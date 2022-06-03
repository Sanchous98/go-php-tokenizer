package tokenizer

import (
	"strings"
)

// PhpMagicKeywords can be hooked to add some extra reserved words
var PhpMagicKeywords = map[string]ItemType{
	"ABSTRACT":        TAbstract,
	"ARRAY":           TArray,
	"AS":              TAs,
	"BREAK":           TBreak,
	"CALLABLE":        TCallable,
	"CASE":            TCase,
	"CATCH":           TCatch,
	"CLASS":           TClass,
	"__CLASS__":       TClassC,
	"CLONE":           TClone,
	"CONST":           TConst,
	"CONTINUE":        TContinue,
	"DECLARE":         TDeclare,
	"DEFAULT":         TDefault,
	"__DIR__":         TDir,
	"DO":              TDo,
	"ECHO":            TEcho,
	"ELSE":            TElse,
	"ELSEIF":          TElseif,
	"EMPTY":           TEmpty,
	"ENDDECLARE":      TEnddeclare,
	"ENDFOR":          TEndfor,
	"ENDFOREACH":      TEndforeach,
	"ENDIF":           TEndif,
	"ENDSWITCH":       TEndswitch,
	"ENDVAL":          TEndwhile,
	"EVAL":            TEval,
	"EXIT":            TExit,
	"DIE":             TExit,
	"EXTENDS":         TExtends,
	"__FILE__":        TFile,
	"FINAL":           TFinal,
	"FINALLY":         TFinally,
	"FOR":             TFor,
	"FOREACH":         TForeach,
	"FUNCTION":        TFunction,
	"CFUNCTION":       TFunction, // ?
	"__FUNCTION__":    TFuncC,
	"GLOBAL":          TGlobal,
	"GOTO":            TGoto,
	"__HALT_COMPILER": THaltCompiler,
	"IF":              TIf,
	"IMPLEMENTS":      TImplements,
	"INCLUDE":         TInclude,
	"INCLUDE_ONCE":    TIncludeOnce,
	"INSTANCEOF":      TInstanceof,
	"INSTEADOF":       TInsteadof,
	"INTERFACE":       TInterface,
	"ISSET":           TIsset,
	"__LINE__":        TLine,
	"LIST":            TList,
	"AND":             TLogicalAnd,
	"OR":              TLogicalOr,
	"XOR":             TLogicalXor,
	"__METHOD__":      TMethodC,
	"NAMESPACE":       TNamespace,
	"__NAMESPACE__":   TNsC,
	"NEW":             TNew,
	"PRINT":           TPrint,
	"PRIVATE":         TPrivate,
	"PUBLIC":          TPublic,
	"PROTECTED":       TProtected,
	"REQUIRE":         TRequire,
	"REQUIRE_ONCE":    TRequireOnce,
	"RETURN":          TReturn,
	"STATIC":          TStatic,
	"SWITCH":          TSwitch,
	"THROW":           TThrow,
	"TRAIT":           TTrait,
	"__TRAIT__":       TTraitC,
	"TRY":             TTry,
	"UNSET":           TUnset,
	"USE":             TUse,
	"VAR":             TVar,
	"WHILE":           TWhile,
	"YIELD":           TYield,
	"YIELD FROM":      TYieldFrom,
}

func lexVariable(l *Lexer) lexState {
	l.advance(1) // '$' (already confirmed)
	if l.acceptPhpLabel() == "" {
		l.emit(Rune('$'))
		return l.base
	}

	l.emit(TVariable)
	return l.base
}

func labelType(lbl string) ItemType {
	// check for PhpMagicKeywords
	for keyword, itemType := range PhpMagicKeywords {
		if strings.EqualFold(keyword, lbl) {
			return itemType
		}
	}

	return TString
}

func lexStringLabel(l *Lexer) lexState {
	lbl := l.acceptPhpLabel()
	t := labelType(lbl)

	if t == TString {
		switch strings.IndexByte(lbl, '\\') {
		case -1:
		case 0:
			t = TNameFullyQualified
		default:
			t = TNameQualified
		}
	}

	l.emit(t)
	if t == THaltCompiler {
		l.emit(TEof)
		return nil
	}
	return l.base
}
