package greeting

const PREFIX = "hello "

func Message(msg string) string {
	if len(msg) == 0 {
		return PREFIX + "thailand"
	}
	return PREFIX + msg
}
