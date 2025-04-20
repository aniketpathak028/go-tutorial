package main

import (
	"fmt"
	"unicode/utf8"
)

func basics() {
	// 1. constants, variables and data types

	// signed integers (positive and negative numbers)
	var intNum int = 23 // default 32/64 bits based on system arch
	var intSixteen int16 = 32767
	var intThirtyTwo int32 = 2147483647

	// unsigned integers (positive numbers only)
	var uintNum uint = 23 // default 32/64 bits based on system arch
	var uintSixteen uint16 = 32767
	var uintThirtyTwo uint32 = 2147483647
	var uintSixtyFour uint64 = 9223372036854775807

	// floating point numbers
	var floatNum float32 = 3.14                     // default 32 bits
	var floatNum64 float64 = 3.14159265358979323846 // default 64 bits

	// complex numbers
	var complexNum complex64 = 1 + 2i    // default 32 bits
	var complexNum64 complex128 = 1 + 2i // default 64 bits

	// boolean
	var boolTrue bool = true
	var boolFalse bool = false

	// string
	var str string = "Hello World"
	str1 := "Hello" + " " + "World"
	fmt.Println(str1)

	// rune
	var myRune rune = 'a'
	fmt.Println("myRune is: ")
	fmt.Print(myRune)

	// len() gives the size needed to store the string in ASCII
	fmt.Println(len("abc"))
	fmt.Println(len("😀"))
	fmt.Println(utf8.RuneCountInString("😀"))

	// arithmetic operators
	// +, -, *, /, %, ++, --
	// can't perform operations with mixed types
	// should be done by casting

	res := floatNum + float32(intThirtyTwo)
	fmt.Println(res)

	// int division leads to integer
	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// if a varibale is not defined Go assigns a default value to it!
	// ufloat, uint, int, float, rune --> 0
	// bool --> false
	// string --> ""

	// initializing multiple variables
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// := can be used when you want the Go compiler to infer the type
	// in runtime (used as a shorthand instead of var var1= 1)

	// best practices
	// when the assigned value is revealing the type we can use :=
	// when it is not, it is a good practice to defined the type
	// here, it is expected unknownVar to be a string
	// ex: var unknownVar string= foo()

	// constants should be assigned when declared and cannot be changed
	const pi float32 = 3.14

	// can't leave a variable declared and unused
	fmt.Println(intNum)
	fmt.Println(intSixteen)
	fmt.Println(intThirtyTwo)
	fmt.Println(uintNum)
	fmt.Println(uintSixteen)
	fmt.Println(uintThirtyTwo)
	fmt.Println(uintSixtyFour)
	fmt.Println(floatNum)
	fmt.Println(floatNum64)
	fmt.Println(complexNum)
	fmt.Println(complexNum64)
	fmt.Println(boolTrue)
	fmt.Println(boolFalse)
	fmt.Println(str)
}
