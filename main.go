package main

import (
	"fmt"
	"os"
	"os/signal"

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
	m := make(map[string]bool)

	reader.ReadInput(os.Stdin, m)

	for v := range m {
		fmt.Println(v)
	}
}
