package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GOLANG")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	//recieve only ch <-chan
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	//send only ch chan<-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
