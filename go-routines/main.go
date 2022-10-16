package main

import (
	"fmt"
	"net/http"
)

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
		getStatusCode(web)
	}
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(5 * time.Millisecond) // we are weating for the thead to compelete and give respose
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops in endpoint ")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
