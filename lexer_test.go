package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNewLexer(t *testing.T) {
	type args struct {
		i           io.Reader
		filename    string
		lexFunction func(*Lexer) lexState
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewLexer(tt.args.i, tt.args.filename, tt.args.lexFunction), "NewLexer(%v, %v, %v)", tt.args.i, tt.args.filename, tt.args.lexFunction)
		})
	}
}
