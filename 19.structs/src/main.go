// Package
// *******
package main

// Imports
// *******
import "fmt"

// Structs
// *******
// Structs are typed collections of fields
// Useful for grouping data together to form records

// Struct to represent a person
type person struct {
	fname string
	lname string
	age   int
}

// Constructor function for a new person struct
func newPerson(fname string, lname string) *person {
	// We can safely return a pointer to local variable
	// A local variable will survive the scope of the function
	per := person{
		fname: fname,
		lname: lname,
	}
	per.age = 42

	return &per
}

// Functions
// *********
func main() {
	// Creating a new struct
	per := person{"Bob", "Green", 20}
	fmt.Println(per)

	// We can also use named fields
	per2 := person{
		fname: "Alice",
		lname: "Fray",
		age:   30,
	}
	fmt.Println(per2)

	// Omitted fields will be zero-values
	per3 := person{fname: "Fred"}
	fmt.Println(per3)

	// Using & on a struct returns a pointer to that struct value in memory
	ptrPer4 := &person{fname: "Ann", lname: "McArthur", age: 40}
	fmt.Println(ptrPer4)

	// But it is idomatic to use a constructor function instead
	ptrPer5 := newPerson("Jon", "Lennon")
	fmt.Println(ptrPer5)

	// We can access struct fields using the dot notation
	fmt.Printf("%s's age is %v\n", per2.fname, per2.age)

	// We can also use dots with struct pointers
	// The pointers are automatically dereferenced
	fmt.Printf("%s's age is %v\n", ptrPer4.fname, ptrPer4.age)

	// Struct are mutable
	// We can change their fields's values
	fmt.Printf("Before: %s's age is %v\n", ptrPer5.fname, ptrPer5.age)
	ptrPer5.age = 51
	fmt.Printf("After: %s's age is %v\n", ptrPer5.fname, ptrPer5.age)

	// If a struct type is only used for a single value,
	// the value can have an anonymous struct type
	// This technique is commonly used for table-driven tests
	dog := struct {
		name   string
		age    int
		isGood bool
	}{
		name:   "Rex",
		age:    4,
		isGood: true,
	}
	fmt.Println(dog)
}

// FOR WINDOWS:
//  To run:                 go run 19.structs\src\main.go
//  To compile:             go build -o 19.structs\bin\structs.exe 19.structs\src\main.go
//  To run after compile:   .\19.structs\bin\structs.exe
//  Compile + Run:          go build -o 19.structs\bin\structs.exe 19.structs\src\main.go && .\19.structs\bin\structs.exe

// FOR LINUX:
//  To run:                 go run 19.structs/src/main.go
//  To compile:             go build -o 19.structs/bin/structs 19.structs/src/main.go
//  To run after compile:   ./19.structs/bin/structs
//  Compile + Run:          go build -o 19.structs/bin/structs 19.structs/src/main.go && ./19.structs/bin/structs
