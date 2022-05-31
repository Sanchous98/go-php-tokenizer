package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexNumber(t *testing.T) {
	tests := []lexerTestCase{
		{"test decimal integer number", NewLexer(strings.NewReader("123"), "", lexNumber), TLnumber, "123"},
		{"test decimal integer with delimiters number", NewLexer(strings.NewReader("123_456"), "", lexNumber), TLnumber, "123_456"},
		{"test decimal float number", NewLexer(strings.NewReader("123.1"), "", lexNumber), TDnumber, "123.1"},
		{"test decimal exponential float number", NewLexer(strings.NewReader("123e-1"), "", lexNumber), TDnumber, "123e-1"},
		{"test binary number", NewLexer(strings.NewReader("0b101"), "", lexNumber), TLnumber, "0b101"},
		{"test octal number", NewLexer(strings.NewReader("0123"), "", lexNumber), TLnumber, "0123"},
		{"test octal explicit number", NewLexer(strings.NewReader("0o123"), "", lexNumber), TLnumber, "0o123"},
		{"test hexadecimal number", NewLexer(strings.NewReader("0xAF"), "", lexNumber), TLnumber, "0xAF"},
		{"negative test binary number", NewLexer(strings.NewReader("0b102"), "", lexNumber), TLnumber, "0b10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
