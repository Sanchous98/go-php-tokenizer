package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/valyala/bytebufferpool"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = rune(-1)

var HookLexer lexState = lexText

type lexState func(l *Lexer) lexState

type Lexer struct {
	input           *bufio.Reader
	filename        string
	start, position int
	width           int
	items           chan Item
	current         Item
	base            lexState

	inputRst []byte
	output   bytebufferpool.ByteBuffer

	startLine, startChar       int // start line/char
	currentLine, currentChar   int // current line/char
	previousLine, previousChar int // previous line/char (for backup)

	baseStack []lexState
}

func NewLexer(i io.Reader, filename string, lexFunction lexState) *Lexer {
	res := &Lexer{
		input:       bufio.NewReader(i),
		filename:    filename,
		items:       make(chan Item, 16),
		startLine:   1,
		currentLine: 1,
	}

	if lexFunction == nil {
		lexFunction = HookLexer
	}

	go res.run(lexFunction(res))

	return res
}

func (l *Lexer) push(state lexState) {
	l.baseStack = append(l.baseStack, l.base)
	l.base = state
}

func (l *Lexer) pop() {
	l.base = l.baseStack[len(l.baseStack)-1]
	l.baseStack = l.baseStack[:len(l.baseStack)-1]
}

func (l *Lexer) write(s string) (int, error) {
	return l.output.WriteString(s)
}

func (l *Lexer) NextItem() (Item, error) {
	l.current = <-l.items
	if l.current.Type == none {
		l.current = Item{Type: TEof}
	}
	if l.current.Type == itemError {
		return l.Current(), errors.New(l.current.Data)
	}
	return l.Current(), nil
}

func (l *Lexer) Current() Item {
	return l.current
}

func (l *Lexer) hasPrefix(s string) bool {
	return l.peekString(len(s), 0) == s
}

func (l *Lexer) hasPrefixI(s string) bool {
	return strings.ToLower(l.peekString(len(s), 0)) == strings.ToLower(s)
}

func (l *Lexer) run(state lexState) {
	l.push(state)
	for s := l.base; s != nil; {
		s = s(l)
	}
	close(l.items)
}

func (l *Lexer) emit(t ItemType) {
	l.items <- Item{t, l.output.String(), Location{l.filename, l.startLine, l.startChar}}
	l.startLine, l.startChar, l.start = l.currentLine, l.currentChar, l.position
	l.output.Reset()
}

func (l *Lexer) next() rune {
	var r rune
	var err error

	if len(l.inputRst) > 0 {
		r, l.width = utf8.DecodeRune(l.inputRst)
		if l.width == len(l.inputRst) {
			l.inputRst = nil
		} else {
			l.inputRst = l.inputRst[l.width:]
		}
	} else {
		r, l.width, err = l.input.ReadRune()
		if err != nil {
			if err != io.EOF {
				l.error(fmt.Sprintf("%v", err))
			}

			return eof
		}
	}

	l.position += l.width
	l.previousChar = l.currentChar

	if r < utf8.RuneSelf {
		l.output.WriteByte(byte(r))
	} else {
		l.output.WriteString(string(r))
	}
	if r == '\n' {
		l.previousLine = l.currentLine
		l.currentLine += 1
		l.currentChar = 0
	} else {
		l.currentChar += 1 // char counts in characters, not in bytes
	}
	return r
}

func (l *Lexer) ignore() {
	l.start = l.position
	l.startLine, l.startChar = l.currentLine, l.currentChar
	l.output.Reset()
}

func (l *Lexer) Reset() {
	tmp := l.output.Bytes()

	if len(l.inputRst) == 0 {
		l.inputRst = tmp
	} else {
		l.inputRst = append(tmp, l.inputRst...)
	}

	l.output.Reset()
	l.position -= len(tmp)
	l.currentLine, l.currentChar = l.startLine, l.startChar
}

func (l *Lexer) backup() {
	if l.width == 0 {
		return
	}

	// update buffers
	tmp := l.output.Bytes()
	r := tmp[len(tmp)-l.width:]  // removed char
	tmp = tmp[:len(tmp)-l.width] // remove
	l.output.Reset()
	l.output.Write(tmp)

	l.inputRst = append(r, l.inputRst...)

	l.position -= l.width
	l.currentLine, l.currentChar = l.previousLine, l.previousChar
	l.width = 0
}

func (l *Lexer) peek(offset int) rune {
	s := stringToBytes(l.peekString(1, offset))
	if len(s) == 0 {
		return eof
	}
	r, _ := utf8.DecodeRune(s)
	return r
}

func (l *Lexer) peekSkipping(str string) rune {
	var i int

	for {
		symbol := l.peek(i)

		if strings.IndexRune(str, symbol) == -1 {
			return symbol
		}

		i++
	}
}

func (l *Lexer) peekUntil(str string) rune {
	var i int

	for {
		symbol := l.peek(i)

		if strings.IndexRune(str, symbol) >= 0 {
			return symbol
		}

		i++
	}
}

func (l *Lexer) peekAllUntil(str string) string {
	var i int

	for {
		symbol := l.peek(i)

		if strings.IndexRune(str, symbol) >= 0 {
			return l.peekString(i+1, 0)
		}

		i++
	}
}

func (l *Lexer) peekAfterWhitespaces() rune {
	return l.peekSkipping(" \t\r\n")
}

func (l *Lexer) peekString(length, offset int) string {
	var s []byte

	switch {
	case len(l.inputRst) >= length+offset:
		s = l.inputRst
	case len(l.inputRst) == 0:
		s, _ = l.input.Peek(length + offset)
	default:
		s, _ = l.input.Peek(length + offset - len(l.inputRst))
		s = append(l.inputRst, s...)
	}

	if len(s) == 0 {
		return ""
	}

	return bytesToString(s[offset : length+offset])
}

func (l *Lexer) peekAllUntilNotEscaped(str string) string {
	var i int

	for {
		symbol := l.peek(i)

		if strings.IndexRune(str, symbol) >= 0 {
			return l.peekString(i+1, 0)
		}

		if symbol == '\\' {
			i++
		}

		i++
	}
}

func (l *Lexer) peekPhpLabel() string {
	// accept a php label, first char is _ or alpha, next chars are alphanumeric or _
	c := l.peek(0)

	if !(unicode.IsLetter(c) || c == '_' || 0x7f <= c || c == '\\') {
		l.backup()
		// we didn't read a single char
		return ""
	}

	var i int

	for {
		i++
		c = l.peek(i)

		if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || 0x7f <= c || c == '\\') {
			return l.peekString(i, 0)
		}
	}
}

func (l *Lexer) advance(c int) {
	for i := 0; i < c; i++ {
		// we do that for two purposes:
		// 1. correctly skip utf-8 characters
		// 2. detect linebreaks so we count these correctly
		l.next()
	}
}

func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptFixed(s string) bool {
	if !l.hasPrefix(s) {
		return false
	}
	l.advance(len(s))
	return true
}

func (l *Lexer) acceptFixedI(s string) bool {
	if !l.hasPrefixI(s) {
		return false
	}
	l.advance(len(s))
	return true
}

func (l *Lexer) acceptSpace() bool {
	return l.accept(" \t\r\n")
}

func (l *Lexer) acceptSpaces() []byte {
	return l.acceptRun(" \t\r\n")
}

func (l *Lexer) acceptRun(valid string) []byte {
	b := bytebufferpool.Get()
	defer bytebufferpool.Put(b)

	for {
		v := l.next()
		switch {
		case strings.IndexRune(valid, v) == -1:
			l.backup()
			fallthrough
		case v == eof:
			return b.Bytes()
		default:
			if v < utf8.RuneSelf {
				b.WriteByte(byte(v))
			} else {
				b.WriteString(string(v))
			}
		}
	}
}

func (l *Lexer) acceptUntil(s string) {
	for strings.IndexRune(s, l.peek(0)) == -1 {
		l.next()
	}
}

func (l *Lexer) acceptUntilRune(r rune) {
	for l.peek(0) != r {
		l.next()
	}
}

func (l *Lexer) acceptUntilFixed(s string) {
loop:
	for _, symbol := range s {
		c := l.next()
		if c == eof {
			return
		}
		if c != symbol {
			goto loop
		}
	}
}

func (l *Lexer) acceptPhpLabel() string {
	// accept a php label, first char is _ or alpha, next chars are alphanumeric or _
	labelStart := l.output.Len()
	c := l.next()

	if !(unicode.IsLetter(c) || c == '_' || 0x7f <= c || c == '\\') {
		l.backup()
		// we didn't read a single char
		return ""
	}

	for {
		c = l.next()

		if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || 0x7f <= c || c == '\\') {
			l.backup()
			return string(l.output.Bytes()[labelStart:])
		}
	}
}

func (l *Lexer) error(str string) lexState {
	l.items <- Item{itemError, str, Location{l.filename, l.startLine, l.startChar}}
	return nil
}
