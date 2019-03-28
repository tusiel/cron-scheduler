package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"./reader"
	"./scheduler"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan

		appCleanup()
		os.Exit(1)
	}()

	start()
}

func appCleanup() {
	fmt.Println("Shutting down...")
}

func start() {
	var currentHour int
	var currentMinute int

	missingTimeParam := len(os.Args) == 1
	if missingTimeParam {
		var err error

		split := strings.Split(time.Now().Format("15:04"), ":")

		currentHour, err = strconv.Atoi(split[0])
		if err != nil {
			fmt.Print("Unable to parse currentHour\n")
			os.Exit(1)
		}

		currentMinute, err = strconv.Atoi(split[1])
		if err != nil {
			fmt.Print("Unable to parse currentMinute\n")
			os.Exit(1)
		}

	} else {
		timeParam := os.Args[1]

		_, err := time.Parse("15:04", timeParam)
		if err != nil {
			fmt.Print("Unable to parse time argument. Make sure it is in the format HH:MM\n")
			os.Exit(1)
		}

		split := strings.Split(timeParam, ":")

		currentHour, err = strconv.Atoi(split[0])
		if err != nil {
			fmt.Print("Unable to parse currentHour\n")
			os.Exit(1)
		}

		currentMinute, err = strconv.Atoi(split[1])
		if err != nil {
			fmt.Print("Unable to parse currentMinute\n")
			os.Exit(1)
		}
	}

	cronMap := make(map[string]bool)
	reader.ReadInput(os.Stdin, cronMap)

	for _, schedule := range scheduler.ProcessSchedules(cronMap, currentHour, currentMinute) {
		fmt.Println(schedule)
	}
}
