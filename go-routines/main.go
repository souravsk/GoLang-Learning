//SO the basic need to of goroutines is to make the process faster. goroutines are mostly like theards To use goroutines we use "go" keyword on the method name but the problem is the main funtion don't wait for the goroutines to go and compelete the process and come back so we use "wait group" which help as to make the main funtion to wait until the goroutines gets compelited.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usely this are in pointers, sync is go package and within sync we have waitgroup

func main() {
	// go greeter("hello") //what we are doing here is create the thead
	// greeter("World")

	websitelist := []string{
		"https://google.com",
		"https://instagram.com",
		"https://souravsk.github.io",
		"https://fb.com",
		"https://gitlab.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() //This wait always is on the last of the main function, the work of the wait() is to make the main() to wait until it get the Done() signl
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(5 * time.Millisecond) // we are weating for the thead to compelete and give respose
// 		fmt.Println(s)
// 	}
// }

//we run it like this it is going to take time because it going to the each website one by one which is time taking
func getStatusCode(endpoint string) {
	defer wg.Done() //this done() should we run atlast so we use the defer so that it executed at the last, the work of done() to give signal to the wait() that work have been down.
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops in endpoint ")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
