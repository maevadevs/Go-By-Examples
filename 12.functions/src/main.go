// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	// Functions are central pieces in Go
	// After defining them, we can call a function as we would expect
	sum := plus(54, 78)
	mul := multiply3(4, 6, 2)

	fmt.Println("plus(54, 78) =", sum)
	fmt.Println("multiply3(4, 6, 2) =", mul)
}

// Takes 2 integers and returns their sum.
func plus(x int64, y int64) int64 {
	return x + y
}

// With multiple consecutive parameters of the same types,
// we can omit the type name for the like-typed parameters
// up to the final parameter that declares the type

// Takes 3 parameters and returns their multiplication.
func multiply3(x, y, z int64) int64 {
	return x * y * z
}

// FOR WINDOWS:
//	To run:					go run 12.functions\src\main.go
//	To compile:				go build -o 12.functions\bin\functions.exe 12.functions\src\main.go
//	To run after compile:	.\12.functions\bin\functions.exe
//	Compile + Run:			go build -o 12.functions\bin\functions.exe 12.functions\src\main.go && .\12.functions\bin\functions.exe

// FOR LINUX:
//	To run:					go run 12.functions/src/main.go
//	To compile:				go build -o 12.functions/bin/functions 12.functions/src/main.go
//	To run after compile:	./12.functions/bin/functions
//	Compile + Run:			go build -o 12.functions/bin/functions 12.functions/src/main.go && ./12.functions/bin/functions
