package main

import "fmt"

type gasEng struct {
	mpg     uint8
	gallons uint8
}

type elecEng struct {
	mpkwh uint8
	kwh   uint8
}

// helps make the canmakeit function more generic
type engine interface {
	milesleft() uint8
}

func (e elecEng) milesleft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEng) milesleft() uint8 {
	return e.gallons * e.mpg
}

func canmakeit(e engine, miles uint8) {
	if miles <= e.milesleft() {
		fmt.Println("Can make it! :)")
	} else {
		fmt.Println("Cannot make it! :(")
	}
}

func Interface() {

	myEng := gasEng{45, 20}
	myEng2 := elecEng{30, 15}

	canmakeit(myEng, 50)
	canmakeit(myEng2, 50)

}
