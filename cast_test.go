package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexPossibleCast(t *testing.T) {
	tests := []lexerTestCase{
		{
			"test int cast",
			NewLexer(strings.NewReader("(int)"), "", lexPossibleCast),
			[]ItemType{TIntCast},
			[]string{"(int)"},
		},
		{
			"test integer cast",
			NewLexer(strings.NewReader("(integer)"), "", lexPossibleCast),
			[]ItemType{TIntCast},
			[]string{"(integer)"},
		},
		{
			"test bool cast",
			NewLexer(strings.NewReader("(bool)"), "", lexPossibleCast),
			[]ItemType{TBoolCast},
			[]string{"(bool)"},
		},
		{
			"test boolean cast",
			NewLexer(strings.NewReader("(boolean)"), "", lexPossibleCast),
			[]ItemType{TBoolCast},
			[]string{"(boolean)"},
		},
		{
			"test float cast",
			NewLexer(strings.NewReader("(float)"), "", lexPossibleCast),
			[]ItemType{TDoubleCast},
			[]string{"(float)"},
		},
		{
			"test double cast",
			NewLexer(strings.NewReader("(double)"), "", lexPossibleCast),
			[]ItemType{TDoubleCast},
			[]string{
				"(double)"},
		},
		{
			"test real cast",
			NewLexer(strings.NewReader("(real)"), "", lexPossibleCast),
			[]ItemType{TDoubleCast},
			[]string{"(real)"},
		},
		{
			"test string cast",
			NewLexer(strings.NewReader("(string)"), "", lexPossibleCast),
			[]ItemType{TStringCast},
			[]string{"(string)"},
		},
		{
			"test array cast",
			NewLexer(strings.NewReader("(array)"), "", lexPossibleCast),
			[]ItemType{TArrayCast},
			[]string{"(array)"},
		},
		{
			"test object cast",
			NewLexer(strings.NewReader("(object)"), "", lexPossibleCast),
			[]ItemType{TObjectCast},
			[]string{"(object)"},
		},
		{
			"test unset cast", NewLexer(strings.NewReader("(unset)"), "", lexPossibleCast),
			[]ItemType{TUnsetCast},
			[]string{"(unset)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
