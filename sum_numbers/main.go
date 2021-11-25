package main

import (
	"fmt"
)

//This function calculates sum of first 100
//natural numbers
//series by series and pushes to channel
//eg: [0,10,20,30,40,50,60,70,80,90,100]
func sumNumbers(consumer chan<- int) {
	defer close(consumer)
	for i := 0; i < 10; i++ {
		sn := 0
		// calculating sum of series
		for j := 1; j <= 9; j++ {
			sn += ((j * 10) + i)
		}
		if i == 0 {
			sn += 100
		} else {
			sn += i
		}
		// returning sum to channel
		consumer <- sn
	}
}
func main() {
	// unbuffered channel
	consumer := make(chan int, 0)

	// initializing go routine
	go sumNumbers(consumer)

	sum := 0

	// for range loop over channel to pull data
	// will exit once the channel is closed
	for n := range consumer {
		sum += n
	}

	// sum will be printed post channel is closed
	fmt.Println("Sum for 100 natural numbers is ", sum)
}
