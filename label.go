package tokenizer

import (
	"strings"
)

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
	switch strings.ToUpper(lbl) {
	case "ABSTRACT":
		return TAbstract
	case "ARRAY":
		return TArray
	case "AS":
		return TAs
	case "BREAK":
		return TBreak
	case "CALLABLE":
		return TCallable
	case "CASE":
		return TCase
	case "CATCH":
		return TCatch
	case "CLASS":
		return TClass
	case "__CLASS__":
		return TClassC
	case "CLONE":
		return TClone
	case "CONST":
		return TConst
	case "CONTINUE":
		return TContinue
	case "DECLARE":
		return TDeclare
	case "DEFAULT":
		return TDefault
	case "__DIR__":
		return TDir
	case "DO":
		return TDo
	case "ECHO":
		return TEcho
	case "ELSE":
		return TElse
	case "ELSEIF":
		return TElseif
	case "EMPTY":
		return TEmpty
	case "ENDDECLARE":
		return TEnddeclare
	case "ENDFOR":
		return TEndfor
	case "ENDFOREACH":
		return TEndforeach
	case "ENDIF":
		return TEndif
	case "ENDSWITCH":
		return TEndswitch
	case "ENDVAL":
		return TEndwhile
	case "EVAL":
		return TEval
	case "EXIT":
		return TExit
	case "DIE":
		return TExit
	case "EXTENDS":
		return TExtends
	case "__FILE__":
		return TFile
	case "FINAL":
		return TFinal
	case "FINALLY":
		return TFinally
	case "FOR":
		return TFor
	case "FOREACH":
		return TForeach
	case "FUNCTION":
		return TFunction
	case "CFUNCTION":
		return TFunction // ?
	case "__FUNCTION__":
		return TFuncC
	case "GLOBAL":
		return TGlobal
	case "GOTO":
		return TGoto
	case "__HALT_COMPILER":
		return THaltCompiler
	case "IF":
		return TIf
	case "IMPLEMENTS":
		return TImplements
	case "INCLUDE":
		return TInclude
	case "INCLUDE_ONCE":
		return TIncludeOnce
	case "INSTANCEOF":
		return TInstanceof
	case "INSTEADOF":
		return TInsteadof
	case "INTERFACE":
		return TInterface
	case "ISSET":
		return TIsset
	case "__LINE__":
		return TLine
	case "LIST":
		return TList
	case "AND":
		return TLogicalAnd
	case "OR":
		return TLogicalOr
	case "XOR":
		return TLogicalXor
	case "__METHOD__":
		return TMethodC
	case "NAMESPACE":
		return TNamespace
	case "__NAMESPACE__":
		return TNsC
	case "NEW":
		return TNew
	case "PRINT":
		return TPrint
	case "PRIVATE":
		return TPrivate
	case "PUBLIC":
		return TPublic
	case "PROTECTED":
		return TProtected
	case "REQUIRE":
		return TRequire
	case "REQUIRE_ONCE":
		return TRequireOnce
	case "RETURN":
		return TReturn
	case "STATIC":
		return TStatic
	case "SWITCH":
		return TSwitch
	case "THROW":
		return TThrow
	case "TRAIT":
		return TTrait
	case "__TRAIT__":
		return TTraitC
	case "TRY":
		return TTry
	case "UNSET":
		return TUnset
	case "USE":
		return TUse
	case "VAR":
		return TVar
	case "WHILE":
		return TWhile
	case "YIELD":
		return TYield
	case "YIELD FROM":
		return TYieldFrom
	default:
		return TString
	}
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
