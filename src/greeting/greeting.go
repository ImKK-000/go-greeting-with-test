package greeting

import (
	"fmt"
)

type Greeting struct {
	Prefix string
}

func (g Greeting) Message(msg string) string {
	if len(msg) == 0 {
		return g.makeStringWithPrefix("thailand")
	}
	return g.makeStringWithPrefix(msg)
}

func (g Greeting) makeStringWithPrefix(msg string) string {
	return fmt.Sprint(g.Prefix, msg)
}
