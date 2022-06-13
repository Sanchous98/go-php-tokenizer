package tokenizer

import (
	"strings"
	"testing"
)

func Test_lexHereDoc(t *testing.T) {
	tests := []lexerTestCase{
		{
			"test heredoc",
			NewLexer(strings.NewReader("<<<PHP\n<?php echo \\$echo;\nPHP;"), "", lexHeredoc),
			[]ItemType{TStartHeredoc, TEncapsedAndWhitespace, TEndHeredoc},
			[]string{"<<<PHP\n", "<?php echo \\$echo;\n", "PHP"},
		},
		{
			"test nowdoc",
			NewLexer(strings.NewReader("<<<'PHP'\n<?php echo $echo;\nPHP;"), "", lexHeredoc),
			[]ItemType{TStartHeredoc, TEncapsedAndWhitespace, TEndHeredoc},
			[]string{"<<<'PHP'\n", "<?php echo $echo;\n", "PHP"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.tokenChecker())
	}
}
