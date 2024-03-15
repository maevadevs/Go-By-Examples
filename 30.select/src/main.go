// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"time"
)

// main
// ****
func main() {
	// Allows to wait on multiple channel operations
	// goroutine + channel + select is very powerful
	// Here, we will select across 2 channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Each channel will receive a value after some amount of time
	// Simulate blocking RPC operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
        // Send to channel 1
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
        // Send to channel 2
		ch2 <- "two"
	}()

    // Using select to await both of these values simultaneously
    // We print each one as it arrives
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("received", msg1)
        case msg2 := <-ch2:
            fmt.Println("received", msg2)
        }
    }
}

// FOR WINDOWS:
//  To run:                 go run 30.select\src\main.go
//  To compile:             go build -o 30.select\bin\select.exe 30.select\src\main.go
//  To run after compile:   .\30.select\bin\select.exe
//  Compile + Run:          go build -o 30.select\bin\select.exe 30.select\src\main.go && .\30.select\bin\select.exe

// FOR LINUX:
//  To run:                 go run 30.select/src/main.go
//  To compile:             go build -o 30.select/bin/select 30.select/src/main.go
//  To run after compile:   ./30.select/bin/select
//  Compile + Run:          go build -o 30.select/bin/select 30.select/src/main.go && ./30.select/bin/select
