package greeting

import (
	"fmt"
)

type Greeting struct {
	Prefix string
}

const DEFAULT_EMPTY_STRING = "thailand"

func (g Greeting) Message(message string) string {
	if len(message) == 0 {
		return g.makeStringWithPrefix(DEFAULT_EMPTY_STRING)
	}
	return g.makeStringWithPrefix(message)
}

func (g Greeting) makeStringWithPrefix(message string) string {
	return fmt.Sprintf("%s %s", g.Prefix, message)
}
