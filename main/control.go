package main

import "fmt"

func control() {
	var result, remainder, err = intDivision(11, 2)

	// error handling and control statements

	// else if and else needs to be in the same line as the closing
	// curly brace
	if err != nil {
		fmt.Print(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer divison is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// switch statement
	switch {
	case err != nil:
		fmt.Print(err.Error())
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
