package tokenizer

import "unicode"

const (
	binaryDigits  = "_01"
	octalDigits   = binaryDigits + "234567"
	decimalDigits = octalDigits + "89"
	hexDigits     = decimalDigits + "abcdefABCDEF"
)

func lexNumber(l *Lexer) lexState {
	// optional leading sign
	l.accept("+-")
	digits := decimalDigits
	var allowDecimal bool
	t := TLnumber

	if l.accept("0") {
		switch {
		case l.accept("xX"):
			digits = hexDigits
		case l.accept("bB"):
			digits = binaryDigits
		case l.accept("oO"), l.peek() != '.':
			digits = octalDigits
		default:
			allowDecimal = true
		}
	} else {
		allowDecimal = true
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

	peek := l.peek()
	// next thing mustn't be alphanumeric
	if unicode.IsLetter(peek) && unicode.IsDigit(peek) {
		l.next()
		return l.error("bad number syntax")
	}
	l.emit(t)
	return l.base
}
