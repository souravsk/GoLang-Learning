package main

import(
	"fmt"
)

func main(){
	fmt.Println("Wellcome to the Struct")
	//This is no inheritance in golang 

	amarjeet := User{"Amarjeet Kumar", "amarjeet@gmail.com", true, 23}
	fmt.Println(amarjeet)
	fmt.Printf("Amarjeet details are %+v\n",amarjeet)
	fmt.Printf("Name is %v and eamil is %v\n", amarjeet.Name, amarjeet.Email)

}

type User struct{
	Name	string
	Email	string
	Status	bool
	Age		int
}