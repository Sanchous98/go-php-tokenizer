package tokenizer

import "strings"

// CastTypes can be hooked to modify types, which can for type cast
var CastTypes = map[string]ItemType{
	"int":     TIntCast,
	"integer": TIntCast,
	"bool":    TBoolCast,
	"boolean": TBoolCast,
	"float":   TDoubleCast,
	"double":  TDoubleCast,
	"real":    TDoubleCast,
	"string":  TStringCast,
	"array":   TArrayCast,
	"object":  TObjectCast,
	"unset":   TUnsetCast,
}

func lexPossibleCast(l *Lexer) lexState {
	// possible (string) etc
	l.next() // "("
	l.acceptSpaces()
	typ := l.acceptPhpLabel()
	l.acceptSpaces()

	if l.accept(")") {
		if token, ok := CastTypes[strings.ToLower(typ)]; ok {
			l.emit(token)
			return l.base
		}
	}

	l.reset() // return to initial state
	l.next()  // "("
	l.emit(Rune('('))

	return l.base
}
