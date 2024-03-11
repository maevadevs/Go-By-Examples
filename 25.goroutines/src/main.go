// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"time"
)

// Functions
// *********
// A goroutine is a lightweight thread of execution
// Allows to execute functions concurrently in an async and non-blocking manner
// Goroutines’ output may be interleaved
// Because goroutines are being run concurrently by the Go runtime
// Whichever function that finish first returns first

// A simple function that runs a short loop.
func simplefunc(from string) {
	for i := range 5 {
		fmt.Println(from, ":", i)
	}
}

// main
// ****

func main() {
	// Here is how we would call a function in the usual way
	// This is running a function synchronously
	simplefunc("Running synchronously")

	// To invoke a go routine, we use the go keyword
	// This new goroutine will execute concurrently with the calling one
	go simplefunc("goroutine async")

	// We can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("goroutine async anonymous")

	// Go routines run separately and async
	// We have to wait for them to finish
	// NOTE: There is a better way to do this using WaitGroup
	time.Sleep(time.Second)
	fmt.Println("-- All done --")

	// NOTE: goroutines’ output may be interleaved
	// Because goroutines are being run concurrently by the Go runtime
	// Whichever function that finish first returns first
	// This is a non-blocking pattern
}

// FOR WINDOWS:
//  To run:                 go run 25.goroutines\src\main.go
//  To compile:             go build -o 25.goroutines\bin\goroutines.exe 25.goroutines\src\main.go
//  To run after compile:   .\25.goroutines\bin\goroutines.exe
//  Compile + Run:          go build -o 25.goroutines\bin\goroutines.exe 25.goroutines\src\main.go && .\25.goroutines\bin\goroutines.exe

// FOR LINUX:
//  To run:                 go run 25.goroutines/src/main.go
//  To compile:             go build -o 25.goroutines/bin/goroutines 25.goroutines/src/main.go
//  To run after compile:   ./25.goroutines/bin/goroutines
//  Compile + Run:          go build -o 25.goroutines/bin/goroutines 25.goroutines/src/main.go && ./25.goroutines/bin/goroutines
