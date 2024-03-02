// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********

// Here we have 2 functions: zeroval and zeroptr

// Arguments are passed to this function by value.
func zeroval(ival int) {
	// ival will be a value-copy, distinct from the passed argument
	ival = 0
}

// Arguments will be passed to this function by pointer reference.
func zeroptr(iptr *int) {
	// iptr param is a pointer to the int value in memory
	// *iptr dereference the pointer from its memory address to the current value at that address
	// Assigning a value to a dereferenced pointer changes the value at the referenced address
	*iptr = 0
}

func main() {
	// Pointers allow to pass references to values
	i := 1
	fmt.Println("initial:", i)

	// Using zeroval does not overwrite the variable
	// Because only the value is passed
	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i, i.e. a pointer to i
	// Using zeroptr overwrites the variable
	// Because we are manipulating the memory location itself
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can also be printed
	// It prints the memory address of the location
	fmt.Println("pointer:", &i)
}

// FOR WINDOWS:
//  To run:                 go run 17.pointers\src\main.go
//  To compile:             go build -o 17.pointers\bin\pointers.exe 17.pointers\src\main.go
//  To run after compile:   .\17.pointers\bin\pointers.exe
//  Compile + Run:          go build -o 17.pointers\bin\pointers.exe 17.pointers\src\main.go && .\17.pointers\bin\pointers.exe

// FOR LINUX:
//  To run:                 go run 17.pointers/src/main.go
//  To compile:             go build -o 17.pointers/bin/pointers 17.pointers/src/main.go
//  To run after compile:   ./17.pointers/bin/pointers
//  Compile + Run:          go build -o 17.pointers/bin/pointers 17.pointers/src/main.go && ./17.pointers/bin/pointers
