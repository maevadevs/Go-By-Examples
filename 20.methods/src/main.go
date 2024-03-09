// Package
// *******
package main

// Imports
// *******
import "fmt"

// Struct
// ******
type rect struct {
	width, height int
}

// Receivers
// *********
// Go support methods defined on struct types
// We call them Receiver in Go
// They can be defined for either pointer or value receiver types
// Go automatically handles conversion between values and pointers for method calls
// We may want to use a pointer receiver type to avoid copying on method calls
// or to allow the method to mutate the receiving struct

// Receiver function to calculate the rectangle's area.
func (r *rect) area() int {
	return r.width * r.height
}

// Receiver function to calculate the rectangle's perimeter.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Functions
// *********
func main() {
	// Go supports methods defined on struct types
	// Initializing a rect struct
	r := rect{width: 10, height: 5}

	// Methods can be called on instance of objects
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

// FOR WINDOWS:
//  To run:                 go run 20.methods\src\main.go
//  To compile:             go build -o 20.methods\bin\structs.exe 20.methods\src\main.go
//  To run after compile:   .\20.methods\bin\structs.exe
//  Compile + Run:          go build -o 20.methods\bin\structs.exe 20.methods\src\main.go && .\20.methods\bin\structs.exe

// FOR LINUX:
//  To run:                 go run 20.methods/src/main.go
//  To compile:             go build -o 20.methods/bin/structs 20.methods/src/main.go
//  To run after compile:   ./20.methods/bin/structs
//  Compile + Run:          go build -o 20.methods/bin/structs 20.methods/src/main.go && ./20.methods/bin/structs
