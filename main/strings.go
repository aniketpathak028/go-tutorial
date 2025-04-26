package main

import (
	"fmt"
	"strings"
)

func Strings() {

	// Strings, Runes and Bytes

	var myString = "résumé"
	var indexed = myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed)

	// prints the underlying byte array in the string containing the
	// bytes for the utf-8 encoding of the string
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// takeaway: when dealing with strings, it's an underlying array of bytes
	// so when you take the length of the string you would get the number
	// of bytes it uses ex:-
	fmt.Println("The length of the string is: ", len(myString))

	// instead cast a string to rune
	// runes are just unicode point numbers that represent the character
	var myRune = []rune("résumé")

	for i, v := range myRune {
		fmt.Println(i, v)
	}

	// and thus we can get the exact length of the rune as
	fmt.Println("The length of the rune is: ", len(myRune))

	// single rune type can be declared
	var myRune1 = 'a'
	fmt.Printf("\n myRune1 = %v", myRune1)

	// strings concat
	// Note: strings are immutable in go meaning they can't be modified
	// once created, so here every time we append a character, a new string
	// get created, to avoid this use strings.Builder
	strSlice := []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	catStr := ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	// string builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}
