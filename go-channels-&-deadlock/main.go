//In goroutines the goroutines function don't know about other gorouttines is there.
//So to deal with the we have "channels". channels are a way in which your multiple goroutines can talk to each other. they will not be aware of what going on other goroutines.
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to go lang channels and Deadlock")

	mychannel := make(chan int, 2) // this how we declear a channel
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// receive only channl
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mychannel)
		wg.Done()
	}(mychannel, wg)

	//send only channal
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mychannel <- 5
		mychannel <- 6
		close(mychannel)
		wg.Done()
	}(mychannel, wg)

	wg.Wait()
}
