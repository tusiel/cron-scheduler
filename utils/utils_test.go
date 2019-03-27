package utils

import "testing"

func TestIsValidInput(t *testing.T) {
	testMap := map[string]bool{
		"* * /folder/directory":  true,
		"1 * /folder/directory":  true,
		"* 15 /folder/directory": true,
		"* /folder/directory":    false,
		"* *":                    false,
	}

	for input, expectedResult := range testMap {
		if IsValidInput(input) != expectedResult {
			t.Errorf("Expected '%s' to have a result of %t but didn't", input, expectedResult)
		}
	}
}
