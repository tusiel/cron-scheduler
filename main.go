package main

import (
	"log"
	"os"
	"os/signal"
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
	log.Println("Shutting down...")
}

func start() {
	log.Println("Starting...")
}
