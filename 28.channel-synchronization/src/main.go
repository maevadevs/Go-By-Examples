// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines
// Example: Using a blocking receive to wait for a goroutine to finish
// NOTE: When waiting for multiple goroutines to finish, prefer using WaitGroup

// Functions
// *********

// Function to run for a goroutine
// done channel - Used to notify another goroutine that this function’s work is done
func worker(done chan bool) {
	// This function does some processing here
	fmt.Print("Working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done!")
	// Notify via the done channel that this function is done
	done <- true
}

// main
// ****
func main() {
	// Start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel
	// NOTE: If we removed <-done, the program would exit before the worker even started
	// We force the program to wait here until the channel notifies back
	<-done
}

// FOR WINDOWS:
//  To run:                 go run 28.channel-synchronization\src\main.go
//  To compile:             go build -o 28.channel-synchronization\bin\channel-synchronization.exe 28.channel-synchronization\src\main.go
//  To run after compile:   .\28.channel-synchronization\bin\channel-synchronization.exe
//  Compile + Run:          go build -o 28.channel-synchronization\bin\channel-synchronization.exe 28.channel-synchronization\src\main.go && .\28.channel-synchronization\bin\channel-synchronization.exe

// FOR LINUX:
//  To run:                 go run 28.channel-synchronization/src/main.go
//  To compile:             go build -o 28.channel-synchronization/bin/channel-synchronization 28.channel-synchronization/src/main.go
//  To run after compile:   ./28.channel-synchronization/bin/channel-synchronization
//  Compile + Run:          go build -o 28.channel-synchronization/bin/channel-synchronization 28.channel-synchronization/src/main.go && ./28.channel-synchronization/bin/channel-synchronization
