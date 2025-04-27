package main

import "fmt"

func generics() {

	// generics can be a useful tool when we want to make a function generic to types

	intSlice := []int8{3, 4, 5}
	floatSlice := []float32{2.1, 3.4, 5.6}

	fmt.Println(sumSlice[int8](intSlice))
	fmt.Println(sumSlice[float32](floatSlice))

}

// can accept any of the types mentioned
// "any" type can be used too in some cases
func sumSlice[T int8 | float32 | int32 | float64 | int64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
