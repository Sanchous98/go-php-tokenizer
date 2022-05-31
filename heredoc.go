package tokenizer

func lexPhpHeredoc(l *Lexer) lexState {
	// we have a string starting with <<<
	if !l.acceptFixed("<<<") {
		l.reset()
		return lexPhpOperator // I guess?
	}
	l.acceptSpaces()

	op := l.acceptPhpLabel()
	if op == "" {
		l.reset()
		return lexPhpOperator
	}

	if !l.accept("\r\n") {
		l.reset()
		return lexPhpOperator
	}

	l.emit(TStartHeredoc)

	op = "\n" + op

	for {
		if l.hasPrefix(op) {
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			l.advance(len(op))
			l.emit(TEndHeredoc)
			break
		}

		c := l.peek()

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
			if l.pos > l.start {
				l.emit(TEncapsedAndWhitespace)
			}
			lexPhpVariable(l) // meh
		default:
			l.next()
		}
	}

	return l.base
}
