package greeting

const PREFIX = "hello "

func Message(msg string) string {
	if len(msg) == 0 {
		return makeStringWithPrefix("thailand")
	}
	return makeStringWithPrefix(msg)
}

func makeStringWithPrefix(msg string) string {
	return PREFIX + msg
}
