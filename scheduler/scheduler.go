package scheduler

import (
	"fmt"
	"strconv"
	"strings"
)

// ProcessSchedules takes a map of cron jobs, the current hour and minute and outputs an array schedules
// with the cron job and when they will next run
func ProcessSchedules(m map[string]bool, currentHour int, currentMinute int) []string {
	schedules := make([]string, 0)

	for v := range m {
		var minute int
		var hour int
		var err error

		split := strings.Split(v, " ")

		if split[0] == "*" {
			minute = currentMinute

			if split[1] != "*" {
				minute = 0
			}
		} else {
			minute, err = strconv.Atoi(split[0])
			if err != nil {
				fmt.Printf("Error converting %s to an integer - %+v", split[0], err)
				continue
			}
		}

		if split[1] == "*" {
			hour = currentHour

			if minute < currentMinute {
				hour++
			}
		} else {
			hour, err = strconv.Atoi(split[1])
			if err != nil {
				fmt.Printf("Error converting %s to an integer - %+v", split[1], err)
				continue
			}
		}

		paddedMinute := fmt.Sprintf("%02d", minute)

		today := fmt.Sprintf("%d:%s Today - %s", hour, paddedMinute, split[2])
		tomorrow := fmt.Sprintf("%d:%s Tomorrow - %s", hour, paddedMinute, split[2])

		var s string

		if hour > currentHour {
			// If we're before the hour we're checking, it's going to be today
			s = today
		} else if hour == currentHour {
			// If the hour is the same as we're comparing, we need to check the minutes
			if currentMinute > minute {
				s = tomorrow
			} else {
				s = today
			}
		} else {
			// If the hour is before now, it's going to be tomorrow
			s = tomorrow
		}

		schedules = append(schedules, s)
	}

	return schedules
}
