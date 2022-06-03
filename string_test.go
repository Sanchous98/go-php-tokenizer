package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexStringConst(t *testing.T) {
	tests := []lexerTestCase{
		{
			"test single quote string",
			NewLexer(strings.NewReader("'teststring'"), "", lexString),
			[]ItemType{TConstantEncapsedString},
			[]string{"'teststring'"},
		},
		{
			"test double quote string",
			NewLexer(strings.NewReader("\"teststring\""), "", lexString),
			[]ItemType{TConstantEncapsedString},
			[]string{"\"teststring\""},
		},
		{
			"test back quote string",
			NewLexer(strings.NewReader("`teststring`"), "", lexString),
			[]ItemType{TConstantEncapsedString},
			[]string{"`teststring`"},
		},
		{
			"test double quote string with implicit interpolation",
			NewLexer(strings.NewReader("\"$teststring\""), "", lexString),
			[]ItemType{Rune('"'), TVariable, Rune('"')},
			[]string{"\"", "$teststring", "\""},
		},
		{
			"test back quote string with implicit interpolation",
			NewLexer(strings.NewReader("`$teststring`"), "", lexString),
			[]ItemType{Rune('`'), TVariable, Rune('`')},
			[]string{"`", "$teststring", "`"},
		},
		{
			"test double quote string with explicit interpolation",
			NewLexer(strings.NewReader("\"{$teststring}\""), "", lexString),
			[]ItemType{Rune('"'), TCurlyOpen, TVariable, Rune('}'), Rune('"')},
			[]string{"\"", "{", "$teststring", "}", "\""},
		},
		{
			"test back quote string with explicit interpolation",
			NewLexer(strings.NewReader("`{$teststring}`"), "", lexString),
			[]ItemType{Rune('`'), TCurlyOpen, TVariable, Rune('}'), Rune('`')},
			[]string{"`", "{", "$teststring", "}", "`"},
		},
		{
			"test double quote string with implicit interpolation with property access",
			NewLexer(strings.NewReader("\"$teststring->property\""), "", lexString),
			[]ItemType{Rune('"'), TVariable, TObjectOperator, TString, Rune('"')},
			[]string{"\"", "$teststring", "->", "property", "\""},
		},
		{
			"test back quote string with implicit interpolation with property access",
			NewLexer(strings.NewReader("`$teststring->property`"), "", lexString),
			[]ItemType{Rune('`'), TVariable, TObjectOperator, TString, Rune('`')},
			[]string{"`", "$teststring", "->", "property", "`"},
		},
		{
			"test double quote string with explicit interpolation with property access",
			NewLexer(strings.NewReader("\"{$teststring->property}\""), "", lexString),
			[]ItemType{Rune('"'), TCurlyOpen, TVariable, TObjectOperator, TString, Rune('}'), Rune('"')},
			[]string{"\"", "{", "$teststring", "->", "property", "}", "\""},
		},
		{
			"test back quote string with explicit interpolation with property access",
			NewLexer(strings.NewReader("`{$teststring->property}`"), "", lexString),
			[]ItemType{Rune('`'), TCurlyOpen, TVariable, TObjectOperator, TString, Rune('}'), Rune('`')},
			[]string{"`", "{", "$teststring", "->", "property", "}", "`"},
		},
		{
			"test double quote string with implicit interpolation with property access and escaped strings",
			NewLexer(strings.NewReader("\" $teststring->property \""), "", lexString),
			[]ItemType{Rune('"'), TEncapsedAndWhitespace, TVariable, TObjectOperator, TString, TEncapsedAndWhitespace, Rune('"')},
			[]string{"\"", " ", "$teststring", "->", "property", " ", "\""},
		},
		{
			"test back quote string with implicit interpolation with property access and escaped strings",
			NewLexer(strings.NewReader("` $teststring->property `"), "", lexString),
			[]ItemType{Rune('`'), TEncapsedAndWhitespace, TVariable, TObjectOperator, TString, TEncapsedAndWhitespace, Rune('`')},
			[]string{"`", " ", "$teststring", "->", "property", " ", "`"},
		},
		{
			"test double quote string with explicit interpolation with property access and escaped strings",
			NewLexer(strings.NewReader("\" {$teststring->property} \""), "", lexString),
			[]ItemType{Rune('"'), TEncapsedAndWhitespace, TCurlyOpen, TVariable, TObjectOperator, TString, Rune('}'), TEncapsedAndWhitespace, Rune('"')},
			[]string{"\"", " ", "{", "$teststring", "->", "property", "}", " ", "\""},
		},
		{
			"test back quote string with explicit interpolation with property access and escaped strings",
			NewLexer(strings.NewReader("` {$teststring->property} `"), "", lexString),
			[]ItemType{Rune('`'), TEncapsedAndWhitespace, TCurlyOpen, TVariable, TObjectOperator, TString, Rune('}'), TEncapsedAndWhitespace, Rune('`')},
			[]string{"`", " ", "{", "$teststring", "->", "property", "}", " ", "`"},
		},
		{
			"test back quote string with explicit interpolation with property and method access and escaped strings",
			NewLexer(strings.NewReader("` {$teststring->property->method()} `"), "", lexString),
			[]ItemType{Rune('`'), TEncapsedAndWhitespace, TCurlyOpen, TVariable, TObjectOperator, TString, TObjectOperator, TString, Rune('('), Rune(')'), Rune('}'), TEncapsedAndWhitespace, Rune('`')},
			[]string{"`", " ", "{", "$teststring", "->", "property", "->", "method", "(", ")", "}", " ", "`"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
