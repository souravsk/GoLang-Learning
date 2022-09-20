package main

import (
	"fmt"
)

func main() {
	fmt.Println("IF else if golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regulate user"
	}else if loginCount > 10{
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 34; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is lees then 10")
	}
	
}