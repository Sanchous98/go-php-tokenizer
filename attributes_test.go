package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexAttributes(t *testing.T) {
	tests := []lexerTestCase{
		{
			"test single attribute",
			NewLexer(strings.NewReader("#[\\Attribute]"), "", lexAttributes),
			[]ItemType{TAttribute, TNameFullyQualified, Rune(']')},
			[]string{"#[", "\\Attribute", "]"},
		},
		{
			"test single attribute",
			NewLexer(strings.NewReader("#[\\Attribute(),\\Attribute]"), "", lexAttributes),
			[]ItemType{TAttribute, TNameFullyQualified, Rune('('), Rune(')'), Rune(','), TNameFullyQualified, Rune(']')},
			[]string{"#[", "\\Attribute", "(", ")", ",", "\\Attribute", "]"},
		},
		{
			"test single attribute",
			NewLexer(strings.NewReader("#[\\Attribute(), \\Attribute]"), "", lexAttributes),
			[]ItemType{TAttribute, TNameFullyQualified, Rune('('), Rune(')'), Rune(','), TWhitespace, TNameFullyQualified, Rune(']')},
			[]string{"#[", "\\Attribute", "(", ")", ",", " ", "\\Attribute", "]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
