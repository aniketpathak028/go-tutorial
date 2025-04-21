package main

import (
	"errors"
	"fmt"
)

// the starting curly brace should be in the declaration line itself

func printMe(printValue string) {
	fmt.Println(printValue)
}

// mention the return type, also possible to return multiple values

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	// note the braces
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
