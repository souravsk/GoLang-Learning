package main

import(
	"fmt"
)

func main(){
	fmt.Println("Welcome to go lang defer\n")
	
	//defer work is when ever the compailer read the defer function to goes to the end of the funtion and the start exquting from there.
	//as you see below ex i want to print hello world but i worte it worng oder even to it will print in a right  way because we use the defer funtion 
	defer fmt.Println("world")
	fmt.Println("Hello")

	// So defer is follow the concept Last in FIrst out LIFO

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	mydefer()
}
	// it's some what like recuretion where value are stored while and after compelting the loops it will print the values
func mydefer(){
	for i := 0; i <5; i++{
		defer fmt.Println(i)
	}
}