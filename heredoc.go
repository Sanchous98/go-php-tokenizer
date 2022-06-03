package tokenizer

func lexHeredoc(l *Lexer) lexState {
	// we have a string starting with <<<
	if !l.acceptFixed("<<<") {
		l.Reset()
		return lexOperator // I guess?
	}
	l.acceptSpaces()

	op := l.acceptPhpLabel()
	if op == "" {
		l.Reset()
		return lexOperator
	}

	if !l.accept("\r\n") {
		l.Reset()
		return lexOperator
	}

	l.emit(TStartHeredoc)

	op = "\n" + op

	for {
		if l.hasPrefix(op) {
			if l.position > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			l.advance(len(op))
			l.emit(TEndHeredoc)
			break
		}

		c := l.peek(0)

		switch c {
		case eof:
			l.emit(TEncapsedAndWhitespace)
			//l.error("unexpected eof in heredoc")
			return nil
		case '\\':
			// advance (ignore) one
			l.next() // \
			l.next() // the escaped char
		case '$':
			// this is a variable
			if l.position > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			lexVariable(l) // meh
		default:
			l.next()
		}
	}

	return l.base
}
