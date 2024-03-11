// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
	// Go only has "for" loops
	// But with different kinds of styles

	// While-style
	fmt.Print("While-Style: ")
	i := 1
	for i <= 5 {
		fmt.Print(i, " ")
		i = i + 1
	}
	fmt.Println()

	// Classic C-Style
	fmt.Print("C-Style: ")
	for j := 1; j <= 5; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()

	// Using range: This is similar to Python-Style
	fmt.Print("Using range: ")
	for k := range 5 {
		fmt.Print(k+1, " ")
	}
	fmt.Println()

	// Infinite Loop + Break
    // Need to make sure that the break condition will be hit at some point
	fmt.Print("Infinite Loop + Break: ")
	l := 1
	for {
		fmt.Print(l, " ")
		l++
		if l > 5 {
			break
		}
	}
	fmt.Println()

	// Skip iteration using continue
	fmt.Print("Even numbers only: ")
	for m := range 10 {
		if m%2 == 1 {
			continue
		}
		fmt.Print(m, " ")
	}
	fmt.Println()
}

// FOR WINDOWS:
//    To run:                   go run 05.for-loops\src\main.go
//    To compile:               go build -o 05.for-loops\bin\for-loops.exe 05.for-loops\src\main.go
//    To run after compile:     .\05.for-loops\bin\for-loops.exe
//    Compile + Run:            go build -o 05.for-loops\bin\for-loops.exe 05.for-loops\src\main.go && .\05.for-loops\bin\for-loops.exe

// FOR LINUX:
//    To run:                   go run 05.for-loops/src/main.go
//    To compile:               go build -o 05.for-loops/bin/for-loops 05.for-loops/src/main.go
//    To run after compile:     ./05.for-loops/bin/for-loops
//    Compile + Run:            go build -o 05.for-loops/bin/for-loops 05.for-loops/src/main.go && ./05.for-loops/bin/for-loops
