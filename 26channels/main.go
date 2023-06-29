package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	// myChannel <- 1
	// fmt.Println(<-myChannel)

	wg.Add(2)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 1
		close(myChannel)
		wg.Done()
	}(myChannel, wg)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <-myChannel
		if isChannelOpen {
			fmt.Println(value)
		}
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
