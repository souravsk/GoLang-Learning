package main

import (
	"fmt"
	// "net/http"
	// "io/ioutil"
	"net/url"
)

const myurl string = "https://lco.dev:3030/lear?cousename=javawithDSA&paymentid=3ec232"

func main(){
	fmt.Println("Welcome to Go lang Web handling")

	// response, err := http.Get(myurl)
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Printf("Response is of type: %T\n",response)
	// defer response.Body.Close() // caller's responsibility to close the connection

	// databytes, err := ioutil.ReadAll(response.Body)
	// if err != nil{
	// 	panic(err)
	// }
	// content := string(databytes)
	// fmt.Println("Body Data:- ",content) //getting the data from the website body 

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of Query Params are: %T \n",qparams)

	fmt.Println(qparams["cousename"])
	fmt.Println(qparams["paymentid"])

	partOfUrl := &url.URL{  //we are changing the url as you can see i use the "&"
		Scheme : "https",
		Host : "github.com",
		Path : "/souravsk",
		RawQuery : "user=souravsk",
	}

	anotherURL := partOfUrl.String() //we are converting the usrl in string using the string funation
	fmt.Println(anotherURL)

}