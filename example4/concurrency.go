package main

// Based on https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3

import (
	"fmt"
)

// Finds the ore
func finder(mine []string, oreChannel chan<- string) { // send only channel
	defer fmt.Println("closing finder go-routine")

	for _, item := range mine { // range over mine till the channel is closed
		if item == "ore" {
			oreChannel <- item //send item on oreChannel
		}
	}
	close(oreChannel)
}

// Ore Breaker
func breaker(oreChannel <-chan string) chan string { // receive only channel
	defer fmt.Println("closing breaker function")

	minedOreChan := make(chan string)

	// spin of
	go func() {
		defer fmt.Println("closing anonymous go-routine in breaker")
		for foundOre := range oreChannel { //read from oreChannel until close
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //send to minedOreChan
		}
		close(minedOreChan)
	}()
	return minedOreChan
}

// Smelts the ore
func smelter(minedOreChan <-chan string, doneChan chan<- bool) {
	defer fmt.Println("closing smelter go-routine")

	for minedOre := range minedOreChan { //read from oreChannel until close
		fmt.Println("From Miner: ", minedOre)
		fmt.Println("From Smelter: Ore is smelted")
	}

	doneChan <- true
}

func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	doneChan := make(chan bool)
	oreChannel := make(chan string)

	go finder(theMine, oreChannel)
	minedOreChan := breaker(oreChannel)
	go smelter(minedOreChan, doneChan)

	// block till done is called, once called the program will run to completion
	<-doneChan
	fmt.Println("bye bye") // printed before the application closes
}
