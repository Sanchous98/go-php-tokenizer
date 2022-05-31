package tokenizer

var lexPhpOps = map[string]ItemType{
	"&=":  TAndEqual,
	"&&":  TBooleanAnd,
	"||":  TBooleanOr,
	"??":  TCoalesce,
	"?>":  TCloseTag,
	".=":  TConcatEqual,
	"--":  TDec,
	"++":  TInc,
	"/=":  TDivEqual,
	"=>":  TDoubleArrow,
	"::":  TPaamayimNekudotayim,
	"...": TEllipsis,
	"==":  TIsEqual,
	">=":  TIsGreaterOrEqual,
	"===": TIsIdentical,
	"!=":  TIsNotEqual,
	"<>":  TIsNotEqual,
	"!==": TIsNotIdentical,
	"<=":  TIsSmallerOrEqual,
	"<=>": TSpaceship,
	"-=":  TMinusEqual,
	"%=":  TModEqual,
	"*=":  TMulEqual,
	"->":  TObjectOperator,
	"|=":  TOrEqual,
	"+=":  TPlusEqual,
	"**":  TPow,
	"**=": TPowEqual,
	"<<":  TSl,
	"<<=": TSlEqual,
	">>":  TSr,
	">>=": TSrEqual,
	"^=":  TXorEqual,
}

func lexPhpOperator(l *Lexer) lexState {
	if t, ok := lexPhpOps[l.peekString(3)]; ok {
		l.advance(3)
		l.emit(t)
		return l.base
	}

	if t, ok := lexPhpOps[l.peekString(2)]; ok {
		l.advance(2)
		if t == TCloseTag {
			// falling back to HTML mode - make linebreak part of closing tag
			l.accept("\r")
			l.accept("\n")
			l.emit(t)
			l.pop()
			return l.base
		}
		l.emit(t)
		return l.base
	}

	l.emit(Rune(l.next()))
	return l.base
}
