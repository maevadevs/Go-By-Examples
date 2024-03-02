// Package
// *******
package main

import (
	"fmt"
)

// Imports
// *******

// Functions
// *********
func main() {
	// Variadic - Functions that accept any number of arguments
	// We can call them with any length of arguments
	var summed int64 = sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Sum: %v\n", summed)

	// We can also call them with slice/array using a spread-like operator
	nums := []float64{189, 251, 327, 492}
	var averaged float64 = avg(nums...)
	fmt.Printf("Avg: %v\n", averaged)
}

// A variadic function to sum an arbitrary number of int arguments.
func sum(nums ...int64) int64 {
	// Here, the type of nums is equivalent to []int64
	var total int64 = 0
	for _, num := range nums {
		total += num
	}
	return total
}

// A variadic function that computes the average of an arbitrary number of arguments.
func avg(nums ...float64) float64 {
	// Here, the type of nums is equivalent to []float64
	var length float64 = float64(len(nums))
	var total float64 = 0
	for _, num := range nums {
		total += num
	}
	return total / length
}

// FOR WINDOWS:
//  To run:                 go run 14.variadic-func\src\main.go
//  To compile:             go build -o 14.variadic-func\bin\variadic-func.exe 14.variadic-func\src\main.go
//  To run after compile:   .\14.variadic-func\bin\variadic-func.exe
//  Compile + Run:          go build -o 14.variadic-func\bin\variadic-func.exe 14.variadic-func\src\main.go && .\14.variadic-func\bin\variadic-func.exe

// FOR LINUX:
//  To run:                 go run 14.variadic-func/src/main.go
//  To compile:             go build -o 14.variadic-func/bin/variadic-func 14.variadic-func/src/main.go
//  To run after compile:   ./14.variadic-func/bin/variadic-func
//  Compile + Run:          go build -o 14.variadic-func/bin/variadic-func 14.variadic-func/src/main.go && ./14.variadic-func/bin/variadic-func
