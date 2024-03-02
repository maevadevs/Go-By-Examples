// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	// Allows to iterate over elements in a variety of structures
	// We can use it with slices and arrays
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	// It provides both the index, value of item from the range
	// We can ignore any of them with _ when not needed
	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	// Using range on map iterates over the Key-Value pairs
	mapping := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range mapping {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range on strings iterates over Unicode code points
	// First value is the starting byte index of the rune
	// Second value is the rune itself
	for bi, runeChar := range "go" {
		fmt.Println(bi, runeChar)
	}
}

// FOR WINDOWS:
//	To run:					go run 11.range\src\main.go
//	To compile:				go build -o 11.range\bin\range.exe 11.range\src\main.go
//	To run after compile:	.\11.range\bin\range.exe
//	Compile + Run:			go build -o 11.range\bin\range.exe 11.range\src\main.go && .\11.range\bin\range.exe

// FOR LINUX:
//	To run:					go run 11.range/src/main.go
//	To compile:				go build -o 11.range/bin/range 11.range/src/main.go
//	To run after compile:	./11.range/bin/range
//	Compile + Run:			go build -o 11.range/bin/range 11.range/src/main.go && ./11.range/bin/range
