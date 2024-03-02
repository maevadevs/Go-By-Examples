// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"slices"
)

// Functions
// *********
func main() {
	// Numbered sequence of elements of a dynamic length
	// Dynamic-length

	// An un-initialized array equals to nil and length 0
	//  This does not allocate memory and the variable points to nil
	var uslice []string
	fmt.Println("Un-initialized slice:", uslice, uslice == nil, len(uslice) == 0)

	// Use `make()` to create an empty slice with non-zero length
	//  By default, the capacity of a new slice is equal to its length
	//  We can pass an initial length as well
	//  This allocates memory and the variable points to memory to a slice with 0 elements
	slc := make([]string, 3)
	fmt.Println("emp:", slc, "len:", len(slc), "cap:", cap(slc))

	// Get and set works just like with arrays
	slc[0] = "a"
	slc[1] = "b"
	slc[2] = "c"
	fmt.Println("set:", slc)
	fmt.Println("get:", slc[2])

	// len(slice) returns the length of the slice
	fmt.Println("len:", len(slc))

	// Slices are richer than array
	// append(): Takes a slice and returns a slice containing additional new values
	// There is no fixed-length as with arrays
	// Slices can grow and shrink dynamically
	slc = append(slc, "d")
	slc = append(slc, "e", "f")
	fmt.Println("apd:", slc, "len:", len(slc), "cap:", cap(slc))

	// We can copy slices using copy()
	newslc := make([]string, len(slc))
	copy(newslc, slc)
	fmt.Println("cpy:", newslc)

	// We can use "slicing" operations: slice[start:end]
	newslc = slc[2:5]
	fmt.Println("sl1:", newslc)

	// We can slice from a min (inclusive)
	newslc = slc[4:]
	fmt.Println("sl2:", newslc)

	// We can slice up to a max (exclusive)
	newslc = slc[:5]
	fmt.Println("sl3:", newslc)

	// Declaring and initializing a slice
	greetings := []string{"Hello", "Holla", "Bonjour"}
	fmt.Println("dcl:", greetings)

	// Use the "slices" pkg for additional utilities
	greets := []string{"Hello", "Holla", "Bonjour"}
	fmt.Println("greetings == greets:", slices.Equal(greetings, greets))

	// We can use slices to create multi-dimensional structures
	// Unlike arrays, lengths can vary for the inner slices
	twoD := make([][]int, 3)

	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// NOTE: For more details about slices, https://go.dev/blog/slices-intro
}

// FOR WINDOWS:
//  To run:                 go run 09.slices\src\main.go
//  To compile:             go build -o 09.slices\bin\slices.exe 09.slices\src\main.go
//  To run after compile:   .\09.slices\bin\slices.exe
//  Compile + Run:          go build -o 09.slices\bin\slices.exe 09.slices\src\main.go && .\09.slices\bin\slices.exe

// FOR LINUX:
//  To run:                 go run 09.slices/src/main.go
//  To compile:             go build -o 09.slices/bin/slices 09.slices/src/main.go
//  To run after compile:   ./09.slices/bin/slices
//  Compile + Run:          go build -o 09.slices/bin/slices 09.slices/src/main.go && ./09.slices/bin/slices
