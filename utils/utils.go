package utils

import "strings"

// IsValidInput checks if the input is valid
func IsValidInput(input string) bool {
	if input == "" || len(strings.Split(input, " ")) != 3 {
		return false
	}

	return true
}
