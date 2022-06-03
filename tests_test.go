package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type lexerTestCase struct {
	name       string
	lexer      *Lexer
	wantTypes  []ItemType
	wantValues []string
}

func (tt *lexerTestCase) tokenChecker() func(*testing.T) {
	return func(t *testing.T) {
		for i := 0; i < len(tt.wantTypes); i++ {
			item, err := tt.lexer.NextItem()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantTypes[i], item.Type)
			assert.Equal(t, tt.wantValues[i], item.Data)
		}
	}
}
