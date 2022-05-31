package tokenizer

import (
	"fmt"
	"github.com/Sanchous98/go-php/core/convert"
	"github.com/Sanchous98/go-php/core/phpv"
	"path"
	"runtime"
)

//go:generate stringer -type=ItemType -linecomment
type ItemType int

const (
	itemError ItemType = 0
	TEof      ItemType = 1
	TThrow    ItemType = iota + 256 // T_THROW
	_
	TInclude                           // T_INCLUDE
	TIncludeOnce                       // T_INCLUDE_ONCE
	TRequire                           // T_REQUIRE
	TRequireOnce                       // T_REQUIRE_ONCE
	TLogicalOr                         // T_LOGICAL_OR
	TLogicalXor                        // T_LOGICAL_XOR
	TLogicalAnd                        // T_LOGICAL_AND
	TPrint                             // T_PRINT
	TYield                             // T_YIELD
	TDoubleArrow                       // T_DOUBLE_ARROW
	TYieldFrom                         // T_YIELD_FROM
	TPlusEqual                         // T_PLUS_EQUAL
	TMinusEqual                        // T_MINUS_EQUAL
	TMulEqual                          // T_MUL_EQUAL
	TDivEqual                          // T_DIV_EQUAL
	TConcatEqual                       // T_CONCAT_EQUAL
	TModEqual                          // T_MOD_EQUAL
	TAndEqual                          // T_AND_EQUAL
	TOrEqual                           // T_OR_EQUAL
	TXorEqual                          // T_XOR_EQUAL
	TSlEqual                           // T_SL_EQUAL
	TSrEqual                           // T_SR_EQUAL
	TPowEqual                          // T_POW_EQUAL
	TCoalesceEqual                     // T_COALESCE_EQUAL
	TCoalesce                          // T_COALESCE
	TBooleanOr                         // T_BOOLEAN_OR
	TBooleanAnd                        // T_BOOLEAN_AND
	TAmpersandNotFollowedByVarOrVarArg // T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
	TAmpersandFollowedByVarOrVarArg    // T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG
	TIsEqual                           // T_IS_EQUAL
	TIsNotEqual                        // T_IS_NOT_EQUAL
	TIsIdentical                       // T_IS_IDENTICAL
	TIsNotIdentical                    // T_IS_NOT_IDENTICAL
	TSpaceship                         // T_SPACESHIP
	TIsSmallerOrEqual                  // T_IS_SMALLER_OR_EQUAL
	TIsGreaterOrEqual                  // T_IS_GREATER_OR_EQUAL
	TSl                                // T_SL
	TSr                                // T_SR
	TInstanceof                        // T_INSTANCEOF
	TIntCast                           // T_INT_CAST
	TDoubleCast                        // T_DOUBLE_CAST
	TStringCast                        // T_STRING_CAST
	TArrayCast                         // T_ARRAY_CAST
	TObjectCast                        // T_OBJECT_CAST
	TBoolCast                          // T_BOOL_CAST
	TUnsetCast                         // T_UNSET_CAST
	TPow                               // T_POW
	TClone                             // T_CLONE
	_
	TElseif                 // T_ELSEIF
	TElse                   // T_ELSE
	TLnumber                // T_LNUMBER
	TDnumber                // T_DNUMBER
	TString                 // T_STRING
	TNameFullyQualified     // T_NAME_FULLY_QUALIFIED
	TNameRelative           // T_NAME_RELATIVE
	TNameQualified          // T_NAME_QUALIFIED
	TVariable               // T_VARIABLE
	TInlineHtml             // T_INLINE_HTML
	TEncapsedAndWhitespace  // T_ENCAPSED_AND_WHITESPACE
	TConstantEncapsedString // T_CONSTANT_ENCAPSED_STRING
	TStringVarname          // T_STRING_VARNAME
	TNumString              // T_NUM_STRING
	TEval                   // T_EVAL
	TNew                    // T_NEW
	TExit                   // T_EXIT
	TIf                     // T_IF
	TEndif                  // T_ENDIF
	TEcho                   // T_ECHO
	TDo                     // T_DO
	TWhile                  // T_WHILE
	TEndwhile               // T_ENDWHILE
	TFor                    // T_FOR
	TEndfor                 // T_ENDFOR
	TForeach                // T_FOREACH
	TEndforeach             // T_ENDFOREACH
	TDeclare                // T_DECLARE
	TEnddeclare             // T_ENDDECLARE
	TAs                     // T_AS
	TSwitch                 // T_SWITCH
	TEndswitch              // T_ENDSWITCH
	TCase                   // T_CASE
	TDefault                // T_DEFAULT
	TMatch                  // T_MATCH
	TBreak                  // T_BREAK
	TContinue               // T_CONTINUE
	TGoto                   // T_GOTO
	TFunction               // T_FUNCTION
	TFn                     // T_FN
	TConst                  // T_CONST
	TReturn                 // T_RETURN
	TTry                    // T_TRY
	TCatch                  // T_CATCH
	TFinally                // T_FINALLY
	TUse                    // T_USE
	TInsteadof              // T_INSTEADOF
	TGlobal                 // T_GLOBAL
	TStatic                 // T_STATIC
	TAbstract               // T_ABSTRACT
	TFinal                  // T_FINAL
	TPrivate                // T_PRIVATE
	TProtected              // T_PROTECTED
	TPublic                 // T_PUBLIC
	TReadonly               // T_READONLY
	TVar                    // T_VAR
	TUnset                  // T_UNSET
	TIsset                  // T_ISSET
	TEmpty                  // T_EMPTY
	THaltCompiler           // T_HALT_COMPILER
	TClass                  // T_CLASS
	TTrait                  // T_TRAIT
	TInterface              // T_INTERFACE
	TEnum                   // T_ENUM
	TExtends                // T_EXTENDS
	TImplements             // T_IMPLEMENTS
	TNamespace              // T_NAMESPACE
	TList                   // T_LIST
	TArray                  // T_ARRAY
	TCallable               // T_CALLABLE
	TLine                   // T_LINE
	TFile                   // T_FILE
	TDir                    // T_DIR
	TClassC                 // T_CLASS_C
	TTraitC                 // T_TRAIT_C
	TMethodC                // T_METHOD_C
	TFuncC                  // T_FUNC_C
	TNsC                    // T_NS_C
	TAttribute              // T_ATTRIBUTE
	TInc                    // T_INC
	TDec                    // T_DEC
	TObjectOperator         // T_OBJECT_OPERATOR
	TNullSafeObjectOperator // T_NULLSAFE_OBJECT_OPERATOR
	TComment                // T_COMMENT
	TDocComment             // T_DOC_COMMENT
	TOpenTag                // T_OPEN_TAG
	TOpenTagWithEcho        // T_OPEN_TAG_WITH_ECHO
	TCloseTag               // T_CLOSE_TAG
	TWhitespace             // T_WHITESPACE
	TStartHeredoc           // T_START_HEREDOC
	TEndHeredoc             // T_END_HEREDOC
	TDollarOpenCurlyBraces  // T_DOLLAR_OPEN_CURLY_BRACES
	TCurlyOpen              // T_CURLY_OPEN
	TPaamayimNekudotayim    // T_PAAMAYIM_NEKUDOTAYIM
	TNsSeparator            // T_NS_SEPARATOR
	TEllipsis               // T_ELLIPSIS
	ItemMax
)

type Item struct {
	Type       ItemType
	Data       string
	Filename   string
	Line, Char int
}

func (i *Item) Errorf(format string, arg ...any) error {
	e := fmt.Sprintf(format, arg...)
	return fmt.Errorf("%s in %s on line %d", e, i.Filename, i.Line)
}

func (i *Item) String() string {
	return i.Type.Name()
}

func (i ItemType) Name() string {
	if i > ItemMax {
		return convert.RunesToString([]rune{'\'', i.Rune(), '\''})
	}
	return i.String()
}

func (i *Item) Rune() rune {
	return i.Type.Rune()
}

func (i ItemType) Rune() rune {
	if i < ItemMax {
		return rune(0)
	}
	return rune(i - ItemMax)
}

func (i *Item) IsSingle(r rune) bool {
	if i.Type < ItemMax {
		return false
	}
	return i.Type == ItemType(r)+ItemMax
}

func (i *Item) IsExpressionEnd() bool {
	// TCLOSETAG is acceptable to end an expression
	return i.IsSingle(';') || i.Type == TCloseTag
}

func (i *Item) Unexpected() error {
	_, f, l, _ := runtime.Caller(1)
	return i.Errorf("syntax error from %s:%d, unexpected %s", path.Base(f), l, i)
}

func (i *Item) Loc() *phpv.Loc {
	return &phpv.Loc{Filename: i.Filename, Line: i.Line, Char: i.Char}
}

func Rune(r rune) ItemType {
	return ItemType(r) + ItemMax
}
