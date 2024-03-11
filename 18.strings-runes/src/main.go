// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"unicode/utf8"
)

// Functions
// *********
// A Go string is a read-only slice of bytes
// Treated as containers of text encoded in UTF-8
// In Go, the concept of a single Character is called Rune
// Rune - An integer that represents a Unicode code-point
// More Details: https://go.dev/blog/strings

func examineRune(r rune) {
    // Values enclosed in single quotes are rune literals
    // We can compare a rune value to a rune literal directly
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

// main
// ****

func main() {
	// Hello in Thai
	const GREET_THAI = "สวัสดี"
	fmt.Println("GREET_THAI:", GREET_THAI)

	// Strings are equivalent to a slice []byte
	// len() produces the length of the raw bytes stored within
	fmt.Println("Len:", len(GREET_THAI))

	// Indexing produces the raw byte values at each index
	// Hex values of all the bytes that constitute the code points in the string
	fmt.Print("Indexing: ")

	for i := 0; i < len(GREET_THAI); i++ {
		fmt.Printf("%x ", GREET_THAI[i])
	}

	fmt.Println()

	// To actually get the character (Rune) in a string, we can use the `utf8` package
	// The run-time of RuneCountInString() depends on the size of the string
	// It has to decode each UTF-8 rune sequentially
	// Some characters are represented by multiple UTF-8 code points
	fmt.Println("Rune count:", utf8.RuneCountInString(GREET_THAI))

	// range-loop auto-decodes each rune to UTF-8 along with its code-point offset in the string
	for i, runeValue := range GREET_THAI {
		fmt.Printf("%#U starts at code point %d\n", runeValue, i)
	}

	// We can achieve the same operation using utf8.DecodeRuneInString() explicitly
	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(GREET_THAI); i += w {
		runeValue, width := utf8.DecodeRuneInString(GREET_THAI[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
        // This demonstrates passing a rune value to a function
		examineRune(runeValue)
	}
}

// FOR WINDOWS:
//  To run:                 go run 18.strings-runes\src\main.go
//  To compile:             go build -o 18.strings-runes\bin\strings-runes.exe 18.strings-runes\src\main.go
//  To run after compile:   .\18.strings-runes\bin\strings-runes.exe
//  Compile + Run:          go build -o 18.strings-runes\bin\strings-runes.exe 18.strings-runes\src\main.go && .\18.strings-runes\bin\strings-runes.exe

// FOR LINUX:
//  To run:                 go run 18.strings-runes/src/main.go
//  To compile:             go build -o 18.strings-runes/bin/strings-runes 18.strings-runes/src/main.go
//  To run after compile:   ./18.strings-runes/bin/strings-runes
//  Compile + Run:          go build -o 18.strings-runes/bin/strings-runes 18.strings-runes/src/main.go && ./18.strings-runes/bin/strings-runes
