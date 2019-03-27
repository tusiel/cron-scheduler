package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"./reader"
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
	var currentTime string

	if len(os.Args) == 1 {
		currentTime = time.Now().Format("15:04")
	} else {
		t := os.Args[1]

		_, err := time.Parse("15:04", t)
		if err != nil {
			fmt.Print("Unable to parse time argument. Make sure it is in the format HH:MM\n")
			os.Exit(1)
		}

		currentTime = t
	}

	_ = currentTime

	m := make(map[string]bool)
	reader.ReadInput(os.Stdin, m)

	for v := range m {
		fmt.Println(v)
	}
}
