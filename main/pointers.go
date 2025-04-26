package main

import "fmt"

func pointers() {
	var p *int32
	var i int32
	p = &i // p now stores the mem address of i
	fmt.Printf("The value of the pointer p is --> %v \n", p)
	fmt.Printf("The value of i is --> %v \n", *p)
	*p = 5 // changes the value of the mem location (de-ref)
	fmt.Printf("The value of i is --> %v \n", *p)

	// slice copying uses pointer under the hood
	// so when a slice is copied, the changes in 1 slice
	// will reflect in the other

	var slice1 = []int32{1, 2, 3}
	var slice2 = slice1

	slice2[1] = 10

	fmt.Println(slice1)
	fmt.Println(slice2)

	// pointers can be very useful when writing functions
	var thing1 = [3]int8{5, 2, 1}
	fmt.Printf("The memory location of thing1 is: %p \n", &thing1)
	var result = square(&thing1)

	fmt.Println("result: ", result)

}

func square(thing2 *[3]int8) [3]int8 {
	fmt.Printf("The memory location of thing2 is: %p \n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
