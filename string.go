package tokenizer

func lexPhpStringConst(l *Lexer) lexState {
	stType := l.next() // " or '
	if stType == '"' {
		// too lazy to work this out, let's switch to the other lexer
		l.emit(Rune('"'))
		l.push(lexPhpStringWhitespace)
		return l.base
	}
	if stType == '`' {
		l.emit(Rune('`'))
		l.push(lexPhpStringWhitespaceBack)
		return l.base
	}

	for {
		c := l.next()
		if c == stType {
			// end of string
			l.emit(TConstantEncapsedString)
			return l.base
		}

		if c == '\\' {
			// advance (ignore) one
			l.next()
			continue
		}
	}
}

func lexPhpStringWhitespace(l *Lexer) lexState {
	for {
		c := l.peek()

		switch c {
		case eof:
			l.emit(TEncapsedAndWhitespace)
			l.error("unexpected eof in string")
			return nil
		case '"':
			// end of string
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			l.next() // "
			l.emit(Rune(c))
			l.pop() // return to previous context
			return l.base
		case '\\':
			// advance (ignore) one
			l.next() // \
			l.next() // the escaped char
		case '$':
			// this is a variable
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			// meh :(
			return lexPhpVariable
		default:
			l.next()
		}
	}
}

func lexPhpStringWhitespaceBack(l *Lexer) lexState {
	for {
		c := l.peek()

		switch c {
		case eof:
			l.emit(TEncapsedAndWhitespace)
			l.error("unexpected eof in string")
			return nil
		case '`':
			// end of string
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			l.next() // `
			l.emit(Rune('`'))
			l.pop() // return to previous context
			return l.base
		case '\\':
			// advance (ignore) one
			l.next() // \
			l.next() // the escaped char
		case '$':
			// this is a variable
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			// meh :(
			return lexPhpVariable
		default:
			l.next()
		}
	}
}
