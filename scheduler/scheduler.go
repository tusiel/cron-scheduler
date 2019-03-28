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

	var minute int
	var hour int
	var err error

	for schedule := range m {
		split := strings.Split(schedule, " ")

		scheduleMinute := split[0]
		scheduleHour := split[1]

		if scheduleMinute == "*" {
			minute = currentMinute

			// If the hour isn't "*", then we will be running it on the hour exactly.
			if scheduleHour != "*" {
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

			// If we've gone passed the specified minute, we need to wait until the next hour
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

		today := fmt.Sprintf("%d:%02d Today - %s", hour, minute, split[2])
		tomorrow := fmt.Sprintf("%d:%02d Tomorrow - %s", hour, minute, split[2])

		var s string

		switch {
		case hour > currentHour: // If we're before the hour we're checking, it's going to be today
			s = today
		case hour == currentHour: // If we're on the hour, we need to check the minutes
			if currentMinute > minute {
				s = tomorrow
			} else {
				s = today
			}
		default: // If the hour is before now, it's going to be tomorrow
			s = tomorrow
		}

		schedules = append(schedules, s)
	}

	return schedules
}
