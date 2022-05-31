package tokenizer

import "strings"

var phpMagicKeywords = map[string]ItemType{
	"abstract":        TAbstract,
	"array":           TArray,
	"as":              TAs,
	"break":           TBreak,
	"callable":        TCallable,
	"case":            TCase,
	"catch":           TCatch,
	"class":           TClass,
	"__CLASS__":       TClassC,
	"clone":           TClone,
	"const":           TConst,
	"continue":        TContinue,
	"declare":         TDeclare,
	"default":         TDefault,
	"__DIR__":         TDir,
	"do":              TDo,
	"echo":            TEcho,
	"else":            TElse,
	"elseif":          TElseif,
	"empty":           TEmpty,
	"enddeclare":      TEnddeclare,
	"endfor":          TEndfor,
	"endforeach":      TEndforeach,
	"endif":           TEndif,
	"endswitch":       TEndswitch,
	"endwhile":        TEndwhile,
	"eval":            TEval,
	"exit":            TExit,
	"die":             TExit,
	"extends":         TExtends,
	"__FILE__":        TFile,
	"final":           TFinal,
	"finally":         TFinally,
	"for":             TFor,
	"foreach":         TForeach,
	"function":        TFunction,
	"cfunction":       TFunction, // ?
	"__FUNCTION__":    TFuncC,
	"global":          TGlobal,
	"goto":            TGoto,
	"__halt_compiler": THaltCompiler,
	"if":              TIf,
	"implements":      TImplements,
	"include":         TInclude,
	"include_once":    TIncludeOnce,
	"instanceof":      TInstanceof,
	"insteadof":       TInsteadof,
	"interface":       TInterface,
	"isset":           TIsset,
	"__LINE__":        TLine,
	"list":            TList,
	"and":             TLogicalAnd,
	"or":              TLogicalOr,
	"xor":             TLogicalXor,
	"__METHOD__":      TMethodC,
	"namespace":       TNamespace,
	"__NAMESPACE__":   TNsC,
	"new":             TNew,
	"print":           TPrint,
	"private":         TPrivate,
	"public":          TPublic,
	"protected":       TProtected,
	"require":         TRequire,
	"require_once":    TRequireOnce,
	"return":          TReturn,
	"static":          TStatic,
	"switch":          TSwitch,
	"throw":           TThrow,
	"trait":           TTrait,
	"__TRAIT__":       TTraitC,
	"try":             TTry,
	"unset":           TUnset,
	"use":             TUse,
	"var":             TVar,
	"while":           TWhile,
	"yield":           TYield,
	// yield from T_YIELD_FROM TODO special case
}

func lexPhpVariable(l *Lexer) lexState {
	l.advance(1) // '$' (already confirmed)
	if l.acceptPhpLabel() == "" {
		l.emit(Rune('$'))
		return l.base
	}

	l.emit(TVariable)
	return l.base
}

func labelType(lbl string) ItemType {
	// check for phpMagicKeywords
	if v, ok := phpMagicKeywords[strings.ToLower(lbl)]; ok {
		return v
	}
	if v, ok := phpMagicKeywords[lbl]; ok {
		return v
	}
	return TString
}

func lexPhpString(l *Lexer) lexState {
	lbl := l.acceptPhpLabel()
	t := labelType(lbl)

	l.emit(t)
	if t == THaltCompiler {
		l.emit(TEof)
		return nil
	}
	return l.base
}
