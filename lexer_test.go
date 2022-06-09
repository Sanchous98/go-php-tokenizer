package tokenizer

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"testing"
)

func BenchmarkLexer(b *testing.B) {
	testFile, _ := os.Open("testdata/test.php")
	defer testFile.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lexer := NewLexer(testFile, "", nil)

		for {
			item, err := lexer.NextItem()
			if err != nil || item.Type == TEof {
				break
			}
		}

		testFile.Seek(0, io.SeekStart)
	}
}

func TestLexer(t *testing.T) {
	testFile, err := os.Open("testdata/test.php")
	assert.NoError(t, err)
	defer testFile.Close()

	tokensFile, err := os.Open("testdata/test.csv")
	assert.NoError(t, err)
	defer tokensFile.Close()

	tokens := csv.NewReader(tokensFile)
	lexer := NewLexer(testFile, "", nil)

	for {
		if token, e := tokens.Read(); e == nil {
			tokenType, _ := strconv.Atoi(token[0])
			value := token[1]
			position, _ := strconv.Atoi(token[2])
			item, _ := lexer.NextItem()

			if !assert.Equal(t, value, string(item.Data)) {
				break
			}

			if tokenType != 0 || position != 0 {
				if !assert.Equal(t, tokenType, int(item.Type)) {
					break
				}
			}
		} else {
			//
			break
		}
	}
}
