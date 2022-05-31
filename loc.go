package tokenizer

import "fmt"

type Location struct {
	Filename   string
	Line, Char int
}

func (l *Location) String() string {
	return fmt.Sprintf("in %s:%d", l.Filename, l.Line)
}
