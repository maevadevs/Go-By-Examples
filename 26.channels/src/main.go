// Package
// *******
package main

// Imports
// *******
import "fmt"

// main
// ****
// Channels are the pipes that connect concurrent goroutines
// We can send values into channels from one goroutine
// and will receive those values into another goroutine
// Allows to capture computed values from goroutines back into the main thread

func main() {
	// We create a new channel with make(chan type)
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the syntax: channel <- value
	// We can send a value to a channel from any goroutine
	go func() {
		messages <- "Hi there!"
	}()

	// Receive a value from a channel using the syntax: <-channel
	// We can capture or use the value received from a channel
	msg := <-messages
	fmt.Println(msg)

	// NOTE:
	// By default, sending and receiving block until both the sender and receiver are ready
	// This allows us to wait here without additional async-related jargons
}

// FOR WINDOWS:
//  To run:                 go run 26.channels\src\main.go
//  To compile:             go build -o 26.channels\bin\channels.exe 26.channels\src\main.go
//  To run after compile:   .\26.channels\bin\channels.exe
//  Compile + Run:          go build -o 26.channels\bin\channels.exe 26.channels\src\main.go && .\26.channels\bin\channels.exe

// FOR LINUX:
//  To run:                 go run 26.channels/src/main.go
//  To compile:             go build -o 26.channels/bin/channels 26.channels/src/main.go
//  To run after compile:   ./26.channels/bin/channels
//  Compile + Run:          go build -o 26.channels/bin/channels 26.channels/src/main.go && ./26.channels/bin/channels
