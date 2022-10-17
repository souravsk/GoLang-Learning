package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{} // mutex help to the memory so tha other thred don't comein and create a mess so it will lock it for one goroutines until it compelted.

	var score = []int{0}

	wg.Add(3) // add all the goroutine

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("First goroutines")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done() //we giving single that we are done here
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("second goroutines")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Third goroutines")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait() //wating for the all the goroutines to compelete

	fmt.Println(score)
}
