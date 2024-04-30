// Package
// *******
package main

// Imports
// *******
import "fmt"

// When using channels as function parameters, we can specify its direction
// Whether a channel is meant to only send or receive values
// This increases the type-safety of the program

// Functions
// *********

// This function only accepts a channel for sending values
// Trying to receive from this channel would give a compile-time error
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// This function takes 2 channel arguments
//   - pings: Channel for receiving
//   - pongs: Channel for sending
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// main
// ****

func main() {
	// Create 2 channels
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Send a receiving channel to ping
	ping(pings, "passed message")

	// Send a sending and a receiving channel to pong()
	pong(pings, pongs)

	// Print what is received via the pongs channel
	fmt.Println(<-pongs)
}

// FOR WINDOWS:
//  To run:                 go run 29.channel-direction\src\main.go
//  To compile:             go build -o 29.channel-direction\bin\channel-direction.exe 29.channel-direction\src\main.go
//  To run after compile:   .\29.channel-direction\bin\channel-direction.exe
//  Compile + Run:          go build -o 29.channel-direction\bin\channel-direction.exe 29.channel-direction\src\main.go && .\29.channel-direction\bin\channel-direction.exe

// FOR LINUX:
//  To run:                 go run 29.channel-direction/src/main.go
//  To compile:             go build -o 29.channel-direction/bin/channel-direction 29.channel-direction/src/main.go
//  To run after compile:   ./29.channel-direction/bin/channel-direction
//  Compile + Run:          go build -o 29.channel-direction/bin/channel-direction 29.channel-direction/src/main.go && ./29.channel-direction/bin/channel-direction
