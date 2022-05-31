package tokenizer

func lexNumber(l *Lexer) lexState {
	// optional leading sign
	l.accept("+-")
	digits := "0123456789"
	allowDecimal := true
	t := TLnumber
	if l.accept("0") {
		// can be octal or hexa
		if l.accept("xX") {
			// hex
			digits = "0123456789abcdefABCDEF"
			allowDecimal = false
		} else if l.peek() != '.' {
			// octal
			digits = "01234567"
			allowDecimal = false
		}
	}
	l.acceptRun(digits)

	if allowDecimal {
		if l.accept(".") {
			l.acceptRun(digits)
			t = TDnumber
		}
		if l.accept("eE") {
			l.accept("+-")
			l.acceptRun(digits)
			t = TDnumber
		}
	}

	// next thing mustn't be alphanumeric
	if isAlphaNumeric(l.peek()) {
		l.next()
		return l.error("bad number syntax")
	}
	l.emit(t)
	return l.base
}
