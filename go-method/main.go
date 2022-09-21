package main

import(
	"fmt"
)

func main(){
	fmt.Println("Welcome to Methods")
	//So basicly there is not much differents between function and methods 
	//A function just a fucntion which is decleard outside of the main funtion and then it is called from the main or any function
	//A Method is called method when we create it in a stucts or class 
	amarjeet := User{"Amarjeet Kumar", "amarjeet@gmail.com", true, 23}
	fmt.Println(amarjeet)
	fmt.Printf("Amarjeet details are %+v\n",amarjeet)
	fmt.Printf("Name is %v and eamil is %v\n", amarjeet.Name, amarjeet.Email)
	amarjeet.Getstatus()
	//in this case the copy of the varibale is pass so the main value  is not changed
	amarjeet.NewMail()
	fmt.Printf("Name is %v and eamil is %v\n", amarjeet.Name, amarjeet.Email)

}

type User struct{
	Name	string
	Email	string
	Status	bool
	Age		int
}

func (u User) Getstatus(){
	fmt.Println("Are you here",u.Status)
}

func (u User) NewMail(){
	u.Email = "sourav@mailwala.com"
	fmt.Println("Email of this user is :",u.Email)
}