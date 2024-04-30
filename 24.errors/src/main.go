// Package
// *******
package main

// Imports
// *******
import (
	"errors"
	"fmt"
)

// Types
// *****
// It is possible to use custom types as errors by implementing the Error() method on them
// Here’s a variant on the example above that uses a custom type argError to explicitly
// represent an argument error

type argError struct {
	arg  int
	prob string
}

// Receiver Function for argError: argError.Error()
func (err *argError) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.prob)
}

// Functions
// *********
// Errors in Go are an explicit separate return value
// We do not use any Try-Catch blocks
// It is easier to see which function returns an error
// We handle them using the same language constructs employed for any other, non-error tasks
// By convention, errors are the last return value and have type error, a built-in interface

func funcWithError(arg int) (int, error) {
	if arg == 42 {
		// errors.New() constructs a basic error value with the given error message
		return -1, errors.New("unable to work with 42")
	}
	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}

func funcWithError2(arg int) (int, error) {
	if arg == 42 {
		// In this case we use &argError syntax to build a new struct,
		// supplying values for the two fields arg and prob
		return -1, &argError{arg, "unable to work with it"}
	}
	return arg + 3, nil
}

// Main
// ****

func main() {
	// The 2 loops below test out each of our error-returning functions above
	// Note that the use of an inline error check on the if line is a common idiom in Go code
	for _, i := range []int{7, 42} {
		// This is a common Go idiom, though we can separate in 2 lines as well
		if r, err := funcWithError(i); err != nil {
			fmt.Println("funcWithError failed:", err)
		} else {
			fmt.Println("funcWithError worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		// This is a common Go idiom, though we can separate in 2 lines as well
		if r, err := funcWithError2(i); err != nil {
			fmt.Println("funcWithError2 failed:", err)
		} else {
			fmt.Println("funcWithError2 worked:", r)
		}
	}

	// If we want to programmatically use the data in a custom error,
	// we need to get the error as an instance of the custom error type via type assertion
	_, err := funcWithError2(42)

	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	// For more details about error handling in Go
	// https://go.dev/blog/error-handling-and-go
}

// FOR WINDOWS:
//  To run:                 go run 24.errors\src\main.go
//  To compile:             go build -o 24.errors\bin\errors.exe 24.errors\src\main.go
//  To run after compile:   .\24.errors\bin\errors.exe
//  Compile + Run:          go build -o 24.errors\bin\errors.exe 24.errors\src\main.go && .\24.errors\bin\errors.exe

// FOR LINUX:
//  To run:                 go run 24.errors/src/main.go
//  To compile:             go build -o 24.errors/bin/errors 24.errors/src/main.go
//  To run after compile:   ./24.errors/bin/errors
//  Compile + Run:          go build -o 24.errors/bin/errors 24.errors/src/main.go && ./24.errors/bin/errors
