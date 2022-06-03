package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexOperator(t *testing.T) {
	tests := []lexerTestCase{
		{
			"test ampersand followed by var",
			NewLexer(strings.NewReader("&$var"), "", lexOperator),
			[]ItemType{TAmpersandFollowedByVarOrVarArg},
			[]string{"&"},
		},
		{
			"test ampersand not followed by var",
			NewLexer(strings.NewReader("&101"), "", lexOperator),
			[]ItemType{TAmpersandNotFollowedByVarOrVarArg},
			[]string{"&"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
