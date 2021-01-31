package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	response := "Ping"
	for {
		fmt.Println("Foo is sending:", response)
		channel <- response

		msg := <-channel
		if msg == "Pong" {
			fmt.Println("Foo has received:", msg)
		}
	}
}

func bar(channel chan string) {
	response := "Pong"
	for {
		message := <-channel
		fmt.Println("Bar has received:", message)

		if message == "Ping" {
			fmt.Println("Bar is sending:", response)
		}
		channel <- response

	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	channel := make(chan string, 0)
	go foo(channel) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
