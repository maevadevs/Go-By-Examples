// Package
// *******
package main

// Imports
// *******
import "fmt"

// main
// ****
// Default channels are unbuffered
// They only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
// Buffered channels accept a limited number of values without a corresponding receiver for those values

func main() {
    // A channel of strings buffering up to 2 values
    messages := make(chan string, 2)

    // We can send values into the channel without a corresponding concurrent receive
    messages <- "buffered"
    messages <- "channel"

    // We can receive those values later as usual
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

// FOR WINDOWS:
//  To run:                 go run 27.channel-buffering\src\main.go
//  To compile:             go build -o 27.channel-buffering\bin\channel-buffering.exe 27.channel-buffering\src\main.go
//  To run after compile:   .\27.channel-buffering\bin\channel-buffering.exe
//  Compile + Run:          go build -o 27.channel-buffering\bin\channel-buffering.exe 27.channel-buffering\src\main.go && .\27.channel-buffering\bin\channel-buffering.exe

// FOR LINUX:
//  To run:                 go run 27.channel-buffering/src/main.go
//  To compile:             go build -o 27.channel-buffering/bin/channel-buffering 27.channel-buffering/src/main.go
//  To run after compile:   ./27.channel-buffering/bin/channel-buffering
//  Compile + Run:          go build -o 27.channel-buffering/bin/channel-buffering 27.channel-buffering/src/main.go && ./27.channel-buffering/bin/channel-buffering
