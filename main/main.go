package main

// special package name which indicates the package contains
// executable application code ie. can be built into binary and run

import (
	"fmt"

	"github.com/aniketpathak028/go-tutorials/utils"
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
	fmt.Println(utils.Number())
	printMe("random ass string")

	control()
	datastructures()

}
