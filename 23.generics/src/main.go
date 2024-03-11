// Package
// *******
package main

// Imports
// *******
import "fmt"

// Types
// *****
// Starting with version 1.18, Go added support for generics
// Also known as type parameters

// Node is a struct container that can take any type of value and point to another Node.
type Node[T any] struct {
	next *Node[T]
	val  T
}

// List is a singly-linked list with values of any type stored in a Node.
type List[T any] struct {
	head, tail *Node[T]
}

// Receiver Functions
// ******************
// We can define methods on generic types just like we do on regular types
// But we have to keep the type parameters in place
// The type is List[T], not List

// Receiver function to add a new element into the linked list.
func (lst *List[T]) Push(value T) {
	if lst.tail == nil {
		// The list is empty: Start from head
		lst.head = &Node[T]{val: value}
		lst.tail = lst.head
	} else {
		// The list is not empty: Append a new Node at the tail
		lst.tail.next = &Node[T]{val: value}
		lst.tail = lst.tail.next
	}
}

// Receiver function to get all element in the linked list.
func (lst *List[T]) GetAll() []T {
	// A slice of elements of whatever type in the list
	var elems []T

	// Append all the elements of the linked list to the slice
	for el := lst.head; el != nil; el = el.next {
		elems = append(elems, el.val)
	}

	// Return the slice
	return elems
}

// Functions
// *********
// A generic function takes types as parameters

// Generic that takes a map of any type and returns a slice of its keys.
//
// Type parameters:
//   - K : Has `comparable` constraint, meaning that we can compare values of this type with the == and != operators.
//     Required for map keys in Go.
//   - V : Has `any` constraint, meaning that it is not restricted in any way.
//     `any` is same as `interface{}`
func MapKeys[K comparable, V any](m map[K]V) []K {
	// Create a slice with the length of the map
	r := make([]K, 0, len(m))

	// Add the keys from the map
	for k := range m {
		r = append(r, k)
	}

	// Return
	return r
}

// main
// ****

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}

	// When invoking generic functions, we can often rely on type inference
	// We don’t have to specify the types for K and V when calling MapKeys()
	// The compiler can infer them automatically
	fmt.Println("map:", m)
	fmt.Println("map's keys:", MapKeys(m))

	// But we could also specify them explicitly
	// NOTE: The access order of the map is not guaranteed
	fmt.Println("keys v2:", MapKeys[int, string](m))

	// Testing the linked-list data structure
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

// FOR WINDOWS:
//  To run:                 go run 23.generics\src\main.go
//  To compile:             go build -o 23.generics\bin\generics.exe 23.generics\src\main.go
//  To run after compile:   .\23.generics\bin\generics.exe
//  Compile + Run:          go build -o 23.generics\bin\generics.exe 23.generics\src\main.go && .\23.generics\bin\generics.exe

// FOR LINUX:
//  To run:                 go run 23.generics/src/main.go
//  To compile:             go build -o 23.generics/bin/generics 23.generics/src/main.go
//  To run after compile:   ./23.generics/bin/generics
//  Compile + Run:          go build -o 23.generics/bin/generics 23.generics/src/main.go && ./23.generics/bin/generics
