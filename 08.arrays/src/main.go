// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********
func main() {
    // Numbered sequence of elements of a specific length
    // Fixed-length

    // Declaring an array of 5 integers
    // Default zero-value of int32 is used: 0
    // The type of elements and length are both part of the array’s type
    var a [5]int32
    fmt.Println("emp:", a)

    // We can set a value at an index: array[index] = value
    // We can get a value: array[index]
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // The builtin len() returns the length of an array
    fmt.Println("len:", len(a))

    // We can declare and initialize an array in one line
    myInts := [5]int{1, 2, 3, 4, 5}
    fmt.Println("myInts:", myInts)

    // Array types are one-dimensional
    // But we can compose types to build multi-dimensional data structures
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    // NOTE: Arrays appear in the form [v1 v2 v3] when printed with fmt.Println
}

// FOR WINDOWS:
//	To run:					go run 08.arrays\src\main.go
//	To compile:				go build -o 08.arrays\bin\arrays.exe 08.arrays\src\main.go
//	To run after compile:	.\08.arrays\bin\arrays.exe
//	Compile + Run:			go build -o 08.arrays\bin\arrays.exe 08.arrays\src\main.go && .\08.arrays\bin\arrays.exe

// FOR LINUX:
//	To run:					go run 08.arrays/src/main.go
//	To compile:				go build -o 08.arrays/bin/arrays 08.arrays/src/main.go
//	To run after compile:	./08.arrays/bin/arrays
//	Compile + Run:			go build -o 08.arrays/bin/arrays 08.arrays/src/main.go && ./08.arrays/bin/arrays
