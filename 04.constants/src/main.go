// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"math"
)

// Global Constants
// ****************
const PI float64 = 3.1415

// Functions
// *********
func main() {
	// Constants are the same as variable but immutable
	// They can appear anywhere where variables can
	// But we cannot use walrus operator with constants
	// Constants must be declared with their type
	const PROJECT_NAME string = "Constants"
	fmt.Println("Project name is", PROJECT_NAME)

	// We can derive constants from other expressions
	const VALUE = 3e20 / 5000000
	fmt.Println("The calculated VALUE is", VALUE)

	// It is also possible to get a numeric constant value with no specific type
	// Then give the type only later
	const NUM = 6e20 / 900000
	fmt.Println("NUM before type assignment", NUM)
	fmt.Printf("Type of NUM before: %T\n", NUM)

	// A constant is given a type explicitly or by using it in a context
	fmt.Println("math.Sin(NUM)", math.Sin(NUM))
	fmt.Printf("Type of NUM after: %T\n", NUM)

	// Global constants work in the same way as global variables
	fmt.Printf("Global Constant PI: %v\n", PI)
}

// FOR WINDOWS:
//	To run:					go run 04.constants\src\main.go
//	To compile:				go build -o 04.constants\bin\constants.exe 04.constants\src\main.go
//	To run after compile:	.\04.constants\bin\constants.exe
//	Compile + Run:			go build -o 04.constants\bin\constants.exe 04.constants\src\main.go && .\04.constants\bin\constants.exe

// FOR LINUX:
//	To run:					go run 04.constants/src/main.go
//	To compile:				go build -o 04.constants/bin/constants 04.constants/src/main.go
//	To run after compile:	./04.constants/bin/constants
//	Compile + Run:			go build -o 04.constants/bin/constants 04.constants/src/main.go && ./04.constants/bin/constants
