// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
// Go supports anonymous function, which can be used as closures
// We can return a function as a value of another function

// Return a closure function that returns an incrementing integer.
func intSeq() func() int {
	// Declaring a variable
	i := 0

	// This returned anonymous function is an enclosure for i
	// i exists and is accessed in this context created by intSeq()
	// Allows intSeq() to maintain its own copy of i across different executions of the closure
	return func() int {
		i++
		return i
	}
}

// main
// ****

func main() {
	// Variable as a closure function
	var nextIntGen func() int = intSeq()

	fmt.Println("nextIntGen", nextIntGen())
	fmt.Println("nextIntGen", nextIntGen())
	fmt.Println("nextIntGen", nextIntGen())

	// The state of a closure is unique for each instance of the function
	var anotherIntGen func() int = intSeq()

	fmt.Println("anotherIntGen", anotherIntGen())
	fmt.Println("anotherIntGen", anotherIntGen())

	fmt.Println("nextIntGen", nextIntGen())
}

// FOR WINDOWS:
//  To run:                 go run 15.closures\src\main.go
//  To compile:             go build -o 15.closures\bin\closures.exe 15.closures\src\main.go
//  To run after compile:   .\15.closures\bin\closures.exe
//  Compile + Run:          go build -o 15.closures\bin\closures.exe 15.closures\src\main.go && .\15.closures\bin\closures.exe

// FOR LINUX:
//  To run:                 go run 15.closures/src/main.go
//  To compile:             go build -o 15.closures/bin/closures 15.closures/src/main.go
//  To run after compile:   ./15.closures/bin/closures
//  Compile + Run:          go build -o 15.closures/bin/closures 15.closures/src/main.go && ./15.closures/bin/closures
