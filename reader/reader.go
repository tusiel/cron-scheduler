package reader

import (
	"bufio"
	"os"
	"strings"

	"../utils"
)

// ReadInput reads each line in os.File and adds each unique key to the cronJobs map
func ReadInput(f *os.File, cronJobs map[string]bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()

		if strings.ToLower(text) == "end" {
			break
		}

		if !utils.IsValidInput(text) {
			continue
		}

		cronJobs[text] = true
	}
}
