package tokenizer

func lexHeredoc(l *Lexer) lexState {
	// we have a string starting with <<<
	if !l.acceptFixed("<<<") {
		l.Reset()
		return lexOperator // I guess?
	}
	l.acceptSpaces()

	var isNowDoc bool
	var op string

	if l.peek(0) == '\'' {
		isNowDoc = true
		l.next()
		op = l.acceptPhpLabel()
		l.next()
	} else {
		op = l.acceptPhpLabel()
	}

	if op == "" || !l.accept("\r\n") {
		l.Reset()
		return lexOperator
	}

	l.emit(TStartHeredoc)

	for {
		if l.hasPrefix(op) {
			if l.position > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			l.advance(len(op))
			l.emit(TEndHeredoc)
			return l.base
		}

		switch c := l.peek(0); c {
		case eof:
			l.emit(TEncapsedAndWhitespace)
			//l.error("unexpected eof in heredoc")
			return nil
		case '\\':
			// advance (ignore) one
			l.next() // \

			if isNowDoc {
				continue
			}

			l.next() // the escaped char
		case '$':
			if isNowDoc {
				l.next()
				continue
			}
			// this is a variable
			if l.position > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			lexVariable(l)
		default:
			l.next()
		}
	}
}
