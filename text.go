package tokenizer

func lexText(l *Lexer) lexState {
	for {
		if l.hasPrefix("<?") {
			if l.pos > l.start {
				l.emit(TInlineHtml)
			}
			return lexPhpOpen
		}
		if l.next() == eof {
			break
		}
		if l.output.Len() >= 8192 {
			l.emit(TInlineHtml)
		}
	}

	// reached eof
	if l.pos > l.start {
		l.emit(TInlineHtml)
	}
	l.emit(TEof)
	return nil
}

func lexPhpOpen(l *Lexer) lexState {
	l.advance(2)
	if l.peek() == '=' {
		l.next()
		l.emit(TOpenTagWithEcho)
		l.push(lexPhp)
		return l.base
	}
	l.acceptFixedI("php")
	if !l.acceptSpace() {
		return l.error("php tag should be followed by a whitespace")
	}
	l.emit(TOpenTag)
	l.push(lexPhp)
	return l.base
}
