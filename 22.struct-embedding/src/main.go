// Package
// *******
package main

// Imports
// *******
import "fmt"

// Interface
// *********
// This interface is implemented by base type

type IDescriber interface {
	describe() string
}

// Types
// *****
// Go supports embedding of structs and interfaces
// Express a more seamless composition of types
// The methods of the embedded also become methods of the container
// NOTE: Not to be confused with //go:embed directive
//  Go Directive since 1.16
//  Embed files and folders into the application binary

type base struct {
	num int
}

// Receivers
// *********
// base implements IDescriber

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Here, type `container` embeds type `base`
// An embedding is like a field with type without name

type container struct {
	base
	name string
}

// main
// ****

func main() {
	// When creating structs with literals, initialize the embedding explicitly
	cont := container{
		base: base{
			num: 1,
		},
		name: "some name",
	}

	// We can access the embedded's fields directly
	fmt.Printf("cont={num: %v, name: %v}\n", cont.num, cont.name)

	// Or we can also spell out the full name
	fmt.Printf("cont={num: %v, name: %v}\n", cont.base.num, cont.name)

	// The methods of the embedded also become methods of the container
	fmt.Println("describe:", cont.describe())

	// If a struct implements an interface,
	// embedding that struct also make container implement the interface
	// This technique can be used to bestow interface implementations onto other structs
	var d IDescriber = cont
	fmt.Println("Describer:", d.describe())
}

// FOR WINDOWS:
//  To run:                 go run 22.struct-embedding\src\main.go
//  To compile:             go build -o 22.struct-embedding\bin\struct-embedding.exe 22.struct-embedding\src\main.go
//  To run after compile:   .\22.struct-embedding\bin\struct-embedding.exe
//  Compile + Run:          go build -o 22.struct-embedding\bin\struct-embedding.exe 22.struct-embedding\src\main.go && .\22.struct-embedding\bin\struct-embedding.exe

// FOR LINUX:
//  To run:                 go run 22.struct-embedding/src/main.go
//  To compile:             go build -o 22.struct-embedding/bin/struct-embedding 22.struct-embedding/src/main.go
//  To run after compile:   ./22.struct-embedding/bin/struct-embedding
//  Compile + Run:          go build -o 22.struct-embedding/bin/struct-embedding 22.struct-embedding/src/main.go && ./22.struct-embedding/bin/struct-embedding
