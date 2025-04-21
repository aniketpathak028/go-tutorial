package main

// special package name which indicates the package contains
// executable application code ie. can be built into binary and run

import (
	"fmt"
)

/*
  package and modules convention

  --> module -  all the .go files are part of the same module
   are  part of a file tree with a single go.mod file in the
   root directory of the tree.

   ex:-
	.
	├── utils
	│   └── number.go
	├── http
	│   └── client.go
	│   └── server.go
	├── main
	│   └── main.go
	├── go.mod

  --> in Go, one package == one directory ie. all .go files for a package
   should be contained in the same directory, and a directory should
   contain the .go files for one package only, never have .go files with
   different package names in the same directory.

  --> all non-main packages, the directory name that the code lives in
  should be the same as the package name. When choosing a name you
  should pick something that is short, descriptive, lower case and
  ideally one word

*/

// a main package must contain a main function somewhere
// convention - main func should be present inside main.go
func main() {
	fmt.Println("lesGooooo!")

	// files inside the same package can access each other's code
	// basics()

	// use the imported package util's Number() function
	// fmt.Println(utils.Number())

	printMe("random ass string")

	var result, remainder, err = intDivision(11, 2)

	// error handling
	// else if and else needs to be in the same line as the closing
	// curly brace
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer divison is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// switch statement
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// conditional switch statement
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

	// string formatting with Printf, replaces the %v with the
	// coresponding variables
	fmt.Printf("the result of the integer division is %v with remainder %v", result, remainder)
}
