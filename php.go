package tokenizer

func lex(l *Lexer) lexState {
	// let's try to find out what we are dealing with
	for {
		c := l.peek(0)
		switch c {
		case ' ', '\r', '\n', '\t':
			l.acceptRun(" \r\n\t")
			l.emit(TWhitespace)
		case '(':
			return lexPossibleCast
		case ')', ',', '{', '}', ';':
			l.emit(Rune(l.next()))
		case '$':
			return lexVariable
		case '#':
			return lexEolComment
		case '/':
			// check if // or /* (comments)
			if l.hasPrefix("//") {
				return lexEolComment
			}
			if l.hasPrefix("/*") {
				return lexBlockComment
			}
			return lexOperator
		case '*', '+', '-', '&', '|', '^', '?', '>', '=', ':', '!', '@', '[', ']', '%', '~':
			return lexOperator
		case '.':
			v := l.peekString(2, 0)
			if len(v) == 2 && v[1] >= '0' && v[1] <= '9' {
				return lexNumber
			}
			// if immediately followed by a number, this is actually a DNUMBER
			return lexOperator
		case '<':
			if l.hasPrefix("<<<") {
				return lexHeredoc
			}
			return lexOperator
		case '\'', '`', '"':
			return lexString
		case eof:
			l.emit(TEof)
			return nil
		default:
			// check for potential label start
			switch {
			case '0' <= c && c <= '9':
				return lexNumber
			case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z', c == '_', 0x7f <= c, c == '\\':
				return lexStringLabel
			}
			return l.error("unexpected character %c", c)
		}
	}
}
