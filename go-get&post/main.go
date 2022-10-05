package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go lang GET and POST")

	//performGetRequest()

	performPostJsonRequest()

}

func errorhandler(err error) {
	if err != nil {
		panic(err)
	}
}

func performGetRequest() {
	const myurl = "http://localhost:1111/get"

	respone, err := http.Get(myurl)
	errorhandler(err)

	defer respone.Body.Close()

	fmt.Println("Status Code: ", respone.StatusCode)
	fmt.Println("Content Lenght is :", respone.ContentLength)

	content, err := ioutil.ReadAll(respone.Body)
	errorhandler(err)
	fmt.Println(string(content))

}

func performPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with GoLang",
			"Price" : 0,
			"platform":"learncodeonline.in"

		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	errorhandler(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	errorhandler(err)

	fmt.Println(string(content))
}
