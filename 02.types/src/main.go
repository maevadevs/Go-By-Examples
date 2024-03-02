// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	// Strings
	fmt.Println("Hello Go!")

	// Integers
	fmt.Println(fmt.Sprint(12345))
	fmt.Println("1 + 1 =>", 1+1)
	// Integer divisions are truncated
	fmt.Println("2 / 3 =>", 2/3)

	// Floats
	fmt.Println(fmt.Sprint(3.1415))
	// Floats division are not truncated
	fmt.Println("7.0 / 3.0 =>", 7.0/3.0)

	// Booleans
	fmt.Println(fmt.Sprint(true))
	// Boolean operations are as expected
	fmt.Println("true && false =>", true && false)
	fmt.Println("true || false =>", true || false)
	fmt.Println("!true =>", !true)
}

// FOR WINDOWS:
//  To run:                 go run 02.types\src\main.go
//  To compile:             go build -o 02.types\bin\types.exe 02.types\src\main.go
//  To run after compile:   .\02.types\bin\types.exe
//  Compile + Run:          go build -o 02.types\bin\types.exe 02.types\src\main.go && .\02.types\bin\types.exe

// FOR LINUX:
//  To run:                 go run 02.types/src/main.go
//  To compile:             go build -o 02.types/bin/types 02.types/src/main.go
//  To run after compile:   ./02.types/bin/types
//  Compile + Run:          go build -o 02.types/bin/types 02.types/src/main.go && ./02.types/bin/types
