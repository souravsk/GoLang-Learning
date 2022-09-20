package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Pointer")
	var pointer *int
	var variable int = 1000
	fmt.Println("Just create without any valuse and just variable name", pointer)
	// fmt.Println("Just create without any valuse and with * ", *pointer)
	// fmt.Println("Just create without any valuse and with & ", &pointer)

	pointer = &variable
	fmt.Printf("it may show some garbage value %v \n", pointer)
	fmt.Printf("it will show the value in the variable * %v \n", *pointer)
	fmt.Printf("it will show the address of variable  & %v \n", &pointer)
}