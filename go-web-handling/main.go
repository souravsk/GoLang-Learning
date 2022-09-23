package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "https://lco.dev"

func main(){
	fmt.Println("Welcome to Go lang Web handling")

	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Body Data:- ",content) //getting the data from the website body 

}