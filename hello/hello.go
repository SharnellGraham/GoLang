// // This code prints "Hello, World!" to the console
// package main

// import "fmt" // import fmt package for formatting output

// func main() { // main function, entry point of program
//     fmt.Println("Hello, World!") // print out "Hello, World!" to the console
// }




// // This program imports the rsc.io/quote package and prints a quote from it using the fmt package
// package main 

// import (
// 	"fmt" // Package fmt implements formatted I/O with functions analogous to C's printf and scanf
// 	"rsc.io/quote" // Package quote provides quotes from the Go authors and contributors.
// )

// func main() {
// 	fmt.Println(quote.Go()) // Prints a quote from the Go authors and contributors
// }


package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
