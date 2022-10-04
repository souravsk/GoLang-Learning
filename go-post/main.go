package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Go lang POST")

}

func errorhandler(err string) {
	if err != nil {
		panic(err)
	}
}

func performGetRequest() {
	const myurl = "http://localhost:1111/get"

	respone, err := http.Get(myurl)
	errorhandler(err)

}
