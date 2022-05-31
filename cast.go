package tokenizer

import "strings"

func lexPhpPossibleCast(l *Lexer) lexState {
	// possible (string) etc

	l.next() // "("
	l.acceptSpaces()

	typ := l.acceptPhpLabel()

	l.acceptSpaces()
	if l.accept(")") {

		switch strings.ToLower(typ) {
		case "int", "integer":
			l.emit(TIntCast)
			return l.base
		case "bool", "boolean":
			l.emit(TBoolCast)
			return l.base
		case "float", "double", "real":
			l.emit(TDoubleCast)
			return l.base
		case "string":
			l.emit(TStringCast)
			return l.base
		case "array":
			l.emit(TArrayCast)
			return l.base
		case "object":
			l.emit(TObjectCast)
			return l.base
		case "unset":
			l.emit(TUnsetCast)
			return l.base
		}
	}

	l.reset() // return to initial state
	l.next()  // "("
	l.emit(Rune('('))

	return l.base
}
