package tokenizer

const (
	doubleQuote quotes = '"'
	backQuote   quotes = '`'
	singleQuote quotes = '\''
)

type quotes = rune

func lexString(l *Lexer) lexState {
	stType := l.peek(0)

	// assert if is string
	switch stType {
	case singleQuote, doubleQuote, backQuote:
	default:
		return l.base
	}

	notConstant, length := checkIfConstant(string(stType), l)

	if notConstant {
		return lexInterpolatedString
	}

	l.advance(length)
	l.emit(TConstantEncapsedString)
	return l.base
}

func lexInterpolatedString(l *Lexer) lexState {
	stType := l.next()
	l.emit(Rune(stType)) // quotes

	for {
		switch c := l.next(); c {
		case stType:
			l.emit(Rune(stType))
			return l.base
		case '{':
			if l.peek(0) != '$' {
				break
			}

			l.emit(TCurlyOpen) // {
			l.next()           // $
			lexVariable(l)

			if l.peekString(2, 0) == "->" {
				lexObjectAccess(l)
			}

			l.emit(Rune(l.next())) // }
		case '$':
			lexVariable(l)

			if l.peekString(2, 0) == "->" {
				l.advance(2)

				if l.peekPhpLabel() == "" {
					break
				}

				l.emit(TObjectOperator)
				l.acceptPhpLabel()
				l.emit(TString)
			}
		case '\\':
			l.next()
		}

		switch l.peek(0) {
		case '{':
			if l.peek(1) != '$' {
				break
			}

			fallthrough
		case '$', stType:
			if l.output.Len() > 0 {
				l.emit(TEncapsedAndWhitespace)
			}
		}
	}
}

func checkIfConstant(stType string, l *Lexer) (notConstant bool, i int) {
	for i = 1; ; i++ {
		c := l.peek(i)

		if string(c) == stType {
			i++
			return
		}

		if c == '$' || c == '{' && l.peek(i+1) == '$' {
			notConstant = true
		}

		if c == '\\' {
			i++
		}
	}
}

func lexObjectAccess(l *Lexer) {
	l.advance(2) // ->
	l.emit(TObjectOperator)

	for {
		if c := l.peek(0); c == '}' {
			break
		}

		switch {
		case l.peekString(2, 0) == "->":
			l.advance(2) // ->
			l.emit(TObjectOperator)
		case len(l.acceptPhpLabel()) > 0:
			l.emit(TString)
		default:
			l.emit(Rune(l.next()))
		}
	}
}
