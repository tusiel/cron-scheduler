package scheduler

import (
	"testing"
)

func TestProcessSchedules(t *testing.T) {
	testMap := map[string]bool{
		"30 1 /bin/run_me_daily":       true,
		"45 * /bin/run_me_hourly":      true,
		"* * /bin/run_me_every_minute": true,
		"* 19 /bin/run_me_sixty_times": true,
		"09 16 /bin/run_me_custom_1":   true,
		"11 * /bin/run_me_custom_2":    true,
	}

	expectedResults := []string{
		"19:00 Today - /bin/run_me_sixty_times",
		"1:30 Tomorrow - /bin/run_me_daily",
		"16:45 Today - /bin/run_me_hourly",
		"16:10 Today - /bin/run_me_every_minute",
		"16:09 Tomorrow - /bin/run_me_custom_1",
		"16:11 Today - /bin/run_me_custom_2",
	}

	schedules := ProcessSchedules(testMap, 16, 10)

	for _, result := range expectedResults {
		var testResult bool
		for _, schedule := range schedules {
			if result == schedule {
				testResult = true
			}
		}

		if !testResult {
			t.Errorf("Expected %s to have been in the array of schedues, but it was not", result)
		}
	}
}
