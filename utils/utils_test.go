package utils

import "testing"

func TestIsValidInput(t *testing.T) {
	testMap := map[string]bool{
		"* * /folder/directory":           true,
		"1 * /folder/directory":           true,
		"* 15 /folder/directory":          true,
		"54 15 /folder/directory":         true,
		"7 9 /folder/directory":           true,
		"7 9 /folder/directory          ": true,
		"          7 9 /folder/directory": true,
		"* /folder/directory":             false,
		"* *":                             false,
		"a 9 /folder/directory":           false,
		"61 9 /folder/directory":          false,
		"11 25 /folder/directory":         false,
	}

	for input, expectedResult := range testMap {
		if IsValidInput(input) != expectedResult {
			t.Errorf("Expected '%s' to have a result of %t but didn't", input, expectedResult)
		}
	}
}

func TestIsValidMinute(t *testing.T) {
	testMap := map[string]bool{
		"11":  true,
		"21":  true,
		"07":  true,
		"7":   true,
		"59":  true,
		"60":  true,
		"*":   true,
		"51a": false,
		"aa":  false,
		"61":  false,
	}

	for input, expectedResult := range testMap {
		if isValidMinute(input) != expectedResult {
			t.Errorf("Expected '%s' to have a result of %t but didn't", input, expectedResult)
		}
	}
}
func TestIsValidHour(t *testing.T) {
	testMap := map[string]bool{
		"11":  true,
		"21":  true,
		"07":  true,
		"7":   true,
		"24":  true,
		"*":   true,
		"25":  false,
		"51a": false,
		"aa":  false,
		"61":  false,
	}

	for input, expectedResult := range testMap {
		if isValidHour(input) != expectedResult {
			t.Errorf("Expected '%s' to have a result of %t but didn't", input, expectedResult)
		}
	}
}
