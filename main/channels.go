package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func channels() {
	var c = make(chan int)

	// deadlock: when writing to an unbuffered channel the code exec
	// will be blocked until someone reads from it

	// c <- 1
	// var i = <-c
	// fmt.Println(i)

	// channels with go-routine
	go process(c)
	for i := range c {
		fmt.Println(i)
	}

	// usecase for channels
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)

}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}

func process(c chan int) {
	defer close(c) // this line is exec when func exits, to close the chanel so that no other process waits
	for i := 0; i < 5; i++ {
		c <- i
	}
}
