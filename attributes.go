package tokenizer

func lexAttributes(l *Lexer) lexState {
	l.advance(2) // #[
	l.emit(TAttribute)

	return lexCode
}
