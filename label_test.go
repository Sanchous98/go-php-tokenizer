package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexPhpVariable(t *testing.T) {
	tests := []lexerTestCase{
		{"test variable starting with letter", NewLexer(strings.NewReader("$abc = 0;"), "", lexPhpVariable), TVariable, "$abc"},
		{"test variable starting with underscore", NewLexer(strings.NewReader("$_abc = 0;"), "", lexPhpVariable), TVariable, "$_abc"},
		{"test variable containing underscore, letters and numbers", NewLexer(strings.NewReader("$_abc012 = 0;"), "", lexPhpVariable), TVariable, "$_abc012"},
		{"test variable starting with number", NewLexer(strings.NewReader("$012 = 0;"), "", lexPhpVariable), Rune('$'), "$"},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
