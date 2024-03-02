// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"time"
)

// Functions
// *********
func main() {
	// Switch can be used as conditional across many branches
	i := 2
	fmt.Print("Write ", i, " as ")
    
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// We can use comma for condition with the same case
	// We can also provide default value for any no-match
	today := time.Now().Weekday()

	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

    // Switch can be used without condition
    // This is an alternative to if--else-if--else
    // Case expression can also be a non-constant
    tm := time.Now()

    switch{
    case tm.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // Switch can also work on the type instead of the value
    // Useful to find the type of an interface
    whatAmI := func(i interface{}) {
        switch typ := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", typ)
        }
    }

    // Testing the type of the interface
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

// FOR WINDOWS:
//	To run:					go run 07.switch\src\main.go
//	To compile:				go build -o 07.switch\bin\switch.exe 07.switch\src\main.go
//	To run after compile:	.\07.switch\bin\switch.exe
//	Compile + Run:			go build -o 07.switch\bin\switch.exe 07.switch\src\main.go && .\07.switch\bin\switch.exe

// FOR LINUX:
//	To run:					go run 07.switch/src/main.go
//	To compile:				go build -o 07.switch/bin/switch 07.switch/src/main.go
//	To run after compile:	./07.switch/bin/switch
//	Compile + Run:			go build -o 07.switch/bin/switch 07.switch/src/main.go && ./07.switch/bin/switch
