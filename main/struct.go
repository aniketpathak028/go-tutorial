package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   // either create a property of type owner or the owner itself to directly access its properties
	int
}

type electricEngine struct {
	mpkwh uint8
	kwn   uint8
}

type owner struct {
	name string
}

// methods of a struct
// can access other methods of the struct and the properties inside the struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it! :)")
	} else {
		fmt.Println("Cannot make it! :(")
	}
}

func Struct() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 7}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)

	// anonymous struct, can't be reused
	myStruct := struct {
		age  uint8
		name string
	}{25, "Aniket"}
	fmt.Println(myStruct)

	// it's same as creating an object of a class and calling its method
	fmt.Printf("calling method of struct: %v \n", myEngine.milesLeft())
	canMakeIt(myEngine, 10)
}
