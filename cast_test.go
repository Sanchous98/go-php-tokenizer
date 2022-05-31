package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexPossibleCast(t *testing.T) {
	tests := []lexerTestCase{
		{"test int cast", NewLexer(strings.NewReader("(int)"), "", lexPossibleCast), TIntCast, "(int)"},
		{"test integer cast", NewLexer(strings.NewReader("(integer)"), "", lexPossibleCast), TIntCast, "(integer)"},
		{"test bool cast", NewLexer(strings.NewReader("(bool)"), "", lexPossibleCast), TBoolCast, "(bool)"},
		{"test boolean cast", NewLexer(strings.NewReader("(boolean)"), "", lexPossibleCast), TBoolCast, "(boolean)"},
		{"test float cast", NewLexer(strings.NewReader("(float)"), "", lexPossibleCast), TDoubleCast, "(float)"},
		{"test double cast", NewLexer(strings.NewReader("(double)"), "", lexPossibleCast), TDoubleCast, "(double)"},
		{"test real cast", NewLexer(strings.NewReader("(real)"), "", lexPossibleCast), TDoubleCast, "(real)"},
		{"test string cast", NewLexer(strings.NewReader("(string)"), "", lexPossibleCast), TStringCast, "(string)"},
		{"test array cast", NewLexer(strings.NewReader("(array)"), "", lexPossibleCast), TArrayCast, "(array)"},
		{"test object cast", NewLexer(strings.NewReader("(object)"), "", lexPossibleCast), TObjectCast, "(object)"},
		{"test unset cast", NewLexer(strings.NewReader("(unset)"), "", lexPossibleCast), TUnsetCast, "(unset)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tokenChecker()(t)
		})
	}
}
