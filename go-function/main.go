package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to go lang function")
	sayhello()
	// we can't create a function in function

	result := adder(3,5)
	fmt.Println("Result is : ",result)

	prores,mymess := proadder(4,6,67,6)
	fmt.Println("Pro result is :", prores,mymess)
}

func adder(a int, b int) int { // This are called signater which mean we have to give the return type 
	return a+b
}

func proadder(values ...int) (int, string) { // This ... used when we don't know how much values we are geting or how much we can accept. This is also called as variadic func
	total := 0
	for _,val := range values{
		total += val
	} 
	return total, "if you want return two values"
}

func sayhello(){
	fmt.Println("Namstey from golang")
}