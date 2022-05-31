package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type lexerTestCase struct {
	name      string
	lexer     *Lexer
	wantType  ItemType
	wantValue string
}

func (tt *lexerTestCase) tokenChecker() func(*testing.T) {
	return func(t *testing.T) {
		item, err := tt.lexer.NextItem()
		assert.NoError(t, err)
		assert.Equal(t, tt.wantType, item.Type)
		assert.Equal(t, tt.wantValue, item.Data)
	}
}
