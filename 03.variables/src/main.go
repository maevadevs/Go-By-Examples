// Package
// *******
package main

// Imports
// *******
import "fmt"

// Global Variables
// ****************
// This is a global variable
// Walrus operator is not available here
var pi float64 = 3.1415

// Functions
// *********
func main() {
	// Declaring a variable
	// Variables with no initial values are zero-values
	var counter int
	fmt.Println("Counter starts at", counter)

	// Declaring and initializing a variable
	var myVar string = "Initial Value"
	fmt.Println("myVar:", myVar)

	// Declaring multiple variables at once
	// Must be of the same type
	var age, count int = 34, 10
	fmt.Println("Age is", age)
	fmt.Println("Count is", count)

	// Type can be skipped and infered if provided during declaration
    // However, we could loose track of the type of the variable
	isHuman := true
	fmt.Println("Is he human?", isHuman)

	// We can use the walrus operator for variable with infered types
	// := is for declaring and infering at the same time
	// This is preferred, but it is only available inside functions
	firstName, lastName := "John", "Smith"
	fmt.Printf("You are %s %s, correct?\n", firstName, lastName)

	// Global variable
	fmt.Printf("The value of PI is %v\n", pi)
}

// FOR WINDOWS:
//	To run:					go run 03.variables\src\main.go
//	To compile:				go build -o 03.variables\bin\variables.exe 03.variables\src\main.go
//	To run after compile:	.\03.variables\bin\variables.exe
//	Compile + Run:			go build -o 03.variables\bin\variables.exe 03.variables\src\main.go && .\03.variables\bin\variables.exe

// FOR LINUX:
//	To run:					go run 03.variables/src/main.go
//	To compile:				go build -o 03.variables/bin/variables 03.variables/src/main.go
//	To run after compile:	./03.variables/bin/variables
//	Compile + Run:			go build -o 03.variables/bin/variables 03.variables/src/main.go && ./03.variables/bin/variables
