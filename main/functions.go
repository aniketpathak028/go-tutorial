package main

import "fmt"

// the starting curly brace should be in the declaration line itself

func printMe(printValue string) {
	fmt.Println(printValue)
}

// mention the return type, also possible to return multiple values

func intDivision(numerator int, denominator int) (int, int) {
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder
}
