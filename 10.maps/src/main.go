// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"maps"
)

// Functions
// *********
func main() {
	// Built-in associative data types
	// Similar to Dictionary in Python

	// To create an empty map, we use `make(map[key-type]value-type)`
	var person map[string]string = make(map[string]string)

	// We can set new key-value pairs
	person["firstName"] = "John"
	person["lastName"] = "Anonymous"

	// Printing shows all the kew-value pairs
	fmt.Println("person:", person)

	// Getting value
	// If the key does not exist, the zero-value of the value-type is returned
	fmt.Println("firstName:", person["firstName"])
	fmt.Println("age:", person["age"]) // Empty string

	// Built-in len() returns the number of key-value pairs in the map
	fmt.Println("len:", len(person))

	// Built-in delete() removes a key-value pair from the map
	delete(person, "lastName")
	fmt.Println("person:", person)

	// Built-in clear() removes all key-value pairs from the map
	// This makes the map into an empty map
	clear(person)
	fmt.Println("person:", person)

	// Sometimes, we want to distinguish between "missing key" vs "keys with zero values"
	// We can use the 2nd return value from accessing the map for this
	person["firstName"] = "" // This is the zero-value of string
	person["lastName"] = "Anonymous"
	// This is an example of a real key with zero value ""
	firstName, isKey := person["firstName"]
	if isKey {
		fmt.Println("firstName is a real key with value", firstName)
	} else {
		fmt.Println("firstName is not a real key with zero-value", firstName)
	}
	// This is an example of a missing key
	age, isKey := person["age"]
	if isKey {
		fmt.Println("age is a real key with value", age)
	} else {
		fmt.Println("age is not a real key with zero-value", age)
	}

	// Declaring and initializing a new map
	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("newMap:", newMap)

	// Use the "map" pkg for additional utilities
	var anotherMap map[string]int = map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(newMap, anotherMap) {
		fmt.Println("newMap == anotherMap")
	}

	// NOTE: Note that maps appear in the form map[k:v k:v] when printed with fmt.Println()
}

// FOR WINDOWS:
//	To run:					go run 10.maps\src\main.go
//	To compile:				go build -o 10.maps\bin\maps.exe 10.maps\src\main.go
//	To run after compile:	.\10.maps\bin\maps.exe
//	Compile + Run:			go build -o 10.maps\bin\maps.exe 10.maps\src\main.go && .\10.maps\bin\maps.exe

// FOR LINUX:
//	To run:					go run 10.maps/src/main.go
//	To compile:				go build -o 10.maps/bin/maps 10.maps/src/main.go
//	To run after compile:	./10.maps/bin/maps
//	Compile + Run:			go build -o 10.maps/bin/maps 10.maps/src/main.go && ./10.maps/bin/maps
