package list

import (
	"strings"
)

// Words lists words
func Words(message string) []string {
	return strings.Fields(message)
}
