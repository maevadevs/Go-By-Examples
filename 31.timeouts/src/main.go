// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"time"
)

// Important for programs that connect to external resources
// Also for programs that need to bound execution time
// Easy and elegant to implement thanks to channels and select

// main
// ****
func main() {
	// Suppose we are executing an external call that returns its
	// result on a channel ch1 after 2s
	// The channel is buffered, so the send in the goroutine is nonblocking
	// This is a common pattern to prevent goroutine leaks in case the channel is never read

    ch1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "result 1"
	}()

    // The select implementing a timeout
    //   res := <-ch1 awaits the result
    //   <-time.After awaits a value to be sent after the timeout of 1s
    // Since select proceeds with the first receive that is ready,
    // we will take the timeout case if the operation takes more than the allowed 1s

    select {
    // Awaiting the result from the channel
	case res := <-ch1:
		fmt.Println(res)
    // Or, if time exceeds and still no response from ch1, timeout
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

    // Allowing a longer timeout will succeed
    ch2 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "result 2"
    }()

    select {
    // Awaiting the result from the channel
    case res := <-ch2:
        fmt.Println(res)
    // Or, if time exceeds and still no response from ch1, timeout
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}

// FOR WINDOWS:
//  To run:                 go run 31.timeouts\src\main.go
//  To compile:             go build -o 31.timeouts\bin\timeouts.exe 31.timeouts\src\main.go
//  To run after compile:   .\31.timeouts\bin\timeouts.exe
//  Compile + Run:          go build -o 31.timeouts\bin\timeouts.exe 31.timeouts\src\main.go && .\31.timeouts\bin\timeouts.exe

// FOR LINUX:
//  To run:                 go run 31.timeouts/src/main.go
//  To compile:             go build -o 31.timeouts/bin/timeouts 31.timeouts/src/main.go
//  To run after compile:   ./31.timeouts/bin/timeouts
//  Compile + Run:          go build -o 31.timeouts/bin/timeouts 31.timeouts/src/main.go && ./31.timeouts/bin/timeouts
