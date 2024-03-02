// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	// Same as if statement in all other languages
	// Parentheses are not required around the conditions
	// Braces are required around the body
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Same as if-else statements in all other languages
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if -- else if -- else
	// A variable can be declared in the if condition
	// This variable will be available in subsequent statements
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// We can use logical operators as well
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// NOTE: There is no ternary operator in Go
}

// FOR WINDOWS:
//  To run:                 go run 06.if-else\src\main.go
//  To compile:             go build -o 06.if-else\bin\if-else.exe 06.if-else\src\main.go
//  To run after compile:   .\06.if-else\bin\if-else.exe
//  Compile + Run:          go build -o 06.if-else\bin\if-else.exe 06.if-else\src\main.go && .\06.if-else\bin\if-else.exe

// FOR LINUX:
//  To run:                 go run 06.if-else/src/main.go
//  To compile:             go build -o 06.if-else/bin/if-else 06.if-else/src/main.go
//  To run after compile:   ./06.if-else/bin/if-else
//  Compile + Run:          go build -o 06.if-else/bin/if-else 06.if-else/src/main.go && ./06.if-else/bin/if-else
