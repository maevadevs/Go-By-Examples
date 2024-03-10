// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"math"
	"reflect"
)

// Interfaces
// **********
// Named collections of method signatures
// Acts like a "contract" with an object
// Allows to specify what methods an object must implement
// But does not specify how they should be implemented
type IGeometry interface {
	area() float64
	perim() float64
}

// Structs
// *******
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Receiver Functions
// ******************
// To implement an interface in Go, we just need to implement all
// the methods listed in the interface
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Functions
// *********
// If a variable is an interface type, then we can call methods
// that are in the named interface
func measure(g IGeometry) {
	fmt.Printf("Type: %s\n", reflect.TypeOf(g))
	fmt.Println("Value:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	// The circle and rect struct types both implement the IGeometry interface
	// so we can use instances of these structs as arguments to measure()
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

// For more details about Interface:
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

// FOR WINDOWS:
//  To run:                 go run 21.interfaces\src\main.go
//  To compile:             go build -o 21.interfaces\bin\structs.exe 21.interfaces\src\main.go
//  To run after compile:   .\21.interfaces\bin\structs.exe
//  Compile + Run:          go build -o 21.interfaces\bin\structs.exe 21.interfaces\src\main.go && .\21.interfaces\bin\structs.exe

// FOR LINUX:
//  To run:                 go run 21.interfaces/src/main.go
//  To compile:             go build -o 21.interfaces/bin/structs 21.interfaces/src/main.go
//  To run after compile:   ./21.interfaces/bin/structs
//  Compile + Run:          go build -o 21.interfaces/bin/structs 21.interfaces/src/main.go && ./21.interfaces/bin/structs
