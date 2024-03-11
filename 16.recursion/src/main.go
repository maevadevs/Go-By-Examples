// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
// A recursion is a function that calls itself over and over down to a base case

// A recursive function that calls itself over and over until it reaches a base case.
func factorial(n int) int {
	// Base case
	if n == 0 {
		return 1
	}

	// Self-call toward the base case
	return n * factorial(n-1)
}

// main
// ****

func main() {
	// A recursion is a function that calls itself down to a base case
	fmt.Printf("Factrial(10): %v\n", factorial(10))

	// Closures can also be recursive
	// However, it must be declared with var explicitly before it is defined
	var fibonacci func(n int) int

	// Then, we can define the closure function
	fibonacci = func(n int) int {
		// Base case
		if n < 2 {
			return n
		}
		// Self-call toward the base case
		return fibonacci(n-1) + fibonacci(n-2)
	}

	// Calling
	fmt.Printf("fibonacci(20): %v\n", fibonacci(20))
}

// FOR WINDOWS:
//  To run:                 go run 16.recursion\src\main.go
//  To compile:             go build -o 16.recursion\bin\recursion.exe 16.recursion\src\main.go
//  To run after compile:   .\16.recursion\bin\recursion.exe
//  Compile + Run:          go build -o 16.recursion\bin\recursion.exe 16.recursion\src\main.go && .\16.recursion\bin\recursion.exe

// FOR LINUX:
//  To run:                 go run 16.recursion/src/main.go
//  To compile:             go build -o 16.recursion/bin/recursion 16.recursion/src/main.go
//  To run after compile:   ./16.recursion/bin/recursion
//  Compile + Run:          go build -o 16.recursion/bin/recursion 16.recursion/src/main.go && ./16.recursion/bin/recursion
