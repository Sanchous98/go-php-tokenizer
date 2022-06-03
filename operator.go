package tokenizer

var Operators = map[string]ItemType{
	"&=":  TAndEqual,
	"&&":  TBooleanAnd,
	"||":  TBooleanOr,
	"??=": TCoalesceEqual,
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

func lexOperator(l *Lexer) lexState {
	if t, ok := Operators[l.peekString(3, 0)]; ok {
		l.advance(3)
		l.emit(t)
		return l.base
	}

	if t, ok := Operators[l.peekString(2, 0)]; ok {
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

	if t := l.peek(0); t == '&' {
		l.next()
		probableVar := l.peekAfterWhitespaces()

		if probableVar == '$' {
			l.emit(TAmpersandFollowedByVarOrVarArg)
		} else {
			l.emit(TAmpersandNotFollowedByVarOrVarArg)
		}
		return l.base
	}

	l.emit(Rune(l.next()))
	return l.base
}
