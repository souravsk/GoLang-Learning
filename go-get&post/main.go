package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go lang GET and POST")

	//performGetRequest() //This function is calling the get fucntion

	//performPostJsonRequest() //This function is calling the post fucntion

	performPostFormRequest()

}

func errorhandler(err error) { // This is handling the error that we get if any request get faild
	if err != nil {
		panic(err)
	}
}

func performGetRequest() {
	const myurl = "http://localhost:1111/get" // This is our url

	respone, err := http.Get(myurl) // with the help of the http.Get method we can get the url
	errorhandler(err)               //it's handling the error data

	defer respone.Body.Close() //we are using the defer function to close the get method after the excution of function

	fmt.Println("Status Code: ", respone.StatusCode)
	fmt.Println("Content Lenght is :", respone.ContentLength)

	content, err := ioutil.ReadAll(respone.Body) // we reading all the data form the url and storing it into a variable
	errorhandler(err)
	fmt.Println(string(content))

}

func performPostJsonRequest() {
	const myurl = "http://localhost:1111/post" //creating the url for post funciton

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with GoLang",
			"Price" : 0,
			"platform":"learncodeonline.in"

		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody) // sending the json data that we create
	errorhandler(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body) // reading the url data and storing it in a variable
	errorhandler(err)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:1111/postform" //using the postform form the js file

	// creating the formdata

	data := url.Values{}
	data.Add("firstname", "sourav")
	data.Add("lastname", "kumar")
	data.Add("email", "souravk326@gmail.com")

	response, err := http.PostForm(myurl, data)
	errorhandler(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	errorhandler(err)

	fmt.Println(string(content))
}
