package greeting

import (
	"fmt"
)

type Greeting struct {
	Prefix string
}

func (g Greeting) Message(message string) string {
	if len(message) == 0 {
		return g.makeStringWithPrefix("thailand")
	}
	return g.makeStringWithPrefix(message)
}

func (g Greeting) makeStringWithPrefix(message string) string {
	return fmt.Sprintf("%s %s", g.Prefix, message)
}
