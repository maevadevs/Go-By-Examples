// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	var a, b, c int

	// Go functions support multiple return values
	// This is similar to tuple unpacking in Python
	// This is often used to return both value and error
	a, b = vals()
	fmt.Println(a)
	fmt.Println(b)

	// If only one value is needed, we can discard the unneeded ones
	_, c = vals()
	fmt.Println(c)
}

// A function that returns 2 integer values.
func vals() (int, int) {
	return 3, 7
}

// FOR WINDOWS:
//  To run:                 go run 13.multiple-returns\src\main.go
//  To compile:             go build -o 13.multiple-returns\bin\multiple-returns.exe 13.multiple-returns\src\main.go
//  To run after compile:   .\13.multiple-returns\bin\multiple-returns.exe
//  Compile + Run:          go build -o 13.multiple-returns\bin\multiple-returns.exe 13.multiple-returns\src\main.go && .\13.multiple-returns\bin\multiple-returns.exe

// FOR LINUX:
//  To run:                 go run 13.multiple-returns/src/main.go
//  To compile:             go build -o 13.multiple-returns/bin/multiple-returns 13.multiple-returns/src/main.go
//  To run after compile:   ./13.multiple-returns/bin/multiple-returns
//  Compile + Run:          go build -o 13.multiple-returns/bin/multiple-returns 13.multiple-returns/src/main.go && ./13.multiple-returns/bin/multiple-returns
