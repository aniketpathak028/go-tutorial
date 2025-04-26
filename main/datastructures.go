package main

import (
	"fmt"
)

func datastructures() {

	// Arrays
	// 1. Fixed length
	// 2. Same type
	// 3. Indexable
	// 4. Contiguous in Memory

	// creates an int32 type array with 3 elements having default values
	// for type int32 ie. 0
	var intArr [3]int32

	// can initialize while declaring
	intArr1 := [3]int32{1, 2, 3}
	intArr2 := [...]int32{1, 2, 3, 10, 4}
	fmt.Println(intArr1)
	fmt.Println(intArr2)

	// can print memory addresses for each of them
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices - wrappers around arrays to give a more general, powerful
	// and convinient interface to sequence of data

	intSlice := []int32{4, 5, 6}
	// length is the number of elements in the slice, while capacity is
	// the number of memory locations occupied by the slice
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // appends a new number to the slice
	// if the capacity is exceed, it creates a new slice in mem with
	// greater capacity and copies each element and then appends the
	// new value
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // adding by destructuring
	fmt.Println(intSlice)

	// make function - specificy len and cap
	intSlice3 := make([]int32, 3, 8)
	fmt.Print(intSlice3)

	/*  MAPs */
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap1 := map[string]uint8{
		"Adam":  23,
		"Sarah": 45,
	}

	fmt.Println(myMap1["Adam"])

	// The Map returns the default value of the value type incase no key
	// exists, it also returns a second boolean variable to check if the
	// key is present or not
	var age, isPresent = myMap1["Jason"]
	if isPresent {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Printf("Invalid Name")
	}

	// delete a key by reference
	delete(myMap1, "Adam")

	// loops in GO
	for name, age := range myMap1 {
		fmt.Printf("Name:%v, Age:%v\n", name, age)
	}

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// increment and decrement operators
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

}
