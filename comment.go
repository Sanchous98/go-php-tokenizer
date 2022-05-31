package tokenizer

func lexPhpEolComment(l *Lexer) lexState {
	// this is a simple comment going until end of line
	l.acceptUntil("\r\n")
	l.emit(TComment)
	return l.base
}

func lexPhpBlockComment(l *Lexer) lexState {
	t := TComment
	if l.hasPrefix("/**") {
		t = TDocComment
	}

	l.acceptUntilFixed("*/")
	l.emit(t)

	return l.base
}
