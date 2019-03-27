package reader

import (
	"bufio"
	"os"

	"../utils"
)

// ReadInput reads each line in a os.File and adds each unique key to the cronJobs map
func ReadInput(f *os.File, cronJobs map[string]bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()

		if text == "end" {
			break
		}

		if !utils.IsValidInput(text) {
			continue
		}

		cronJobs[text] = true
	}
}
