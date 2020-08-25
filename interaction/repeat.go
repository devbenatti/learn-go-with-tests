package interaction

import "strings"

// Repeat receives a character and returns five repetitions of it
func Repeat(character string, count int) string {
	return strings.Repeat(character, count)
}
