package count

import "strings"

// Words counts words
func Words(message string) int {
	return len(strings.Fields(message))
}
