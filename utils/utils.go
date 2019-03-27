package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidInput checks if the input is valid.
// Validation rules:
//			- Must be made up of 3 sections (minutes, hours, cron-job)
//			- Minutes must be in the in the range of 0-60 inclusive or *
//			- Hours must be in the in the range of 0-24 inclusive or *
func IsValidInput(input string) bool {
	sections := strings.Split(strings.Trim(input, " "), " ")

	validNumSections := len(sections) == 3

	if !validNumSections || !isValidMinute(sections[0]) || !isValidHour(sections[1]) {
		return false
	}

	return true
}

// isValidMinute checks the minute is a number between 0-60 inclusive or *
func isValidMinute(minute string) bool {
	r, err := regexp.Compile("^(?:[1-9]|0[1-9]|[1-5]+?[0-9]+?|60|[*])$")
	if err != nil {
		fmt.Printf("Error compiling Minute regex: %+v", err)
		return false
	}

	return r.MatchString(minute)
}

// isValidHour checks the minute is a number between 0-24 inclusive or *
func isValidHour(hour string) bool {
	r, err := regexp.Compile("^(?:[1-9]|0[1-9]|1+?[0-9]+?|2+?[0-4]+?|[*])$")
	if err != nil {
		fmt.Printf("Error compiling Hour regex: %+v", err)
		return false
	}

	return r.MatchString(hour)
}
