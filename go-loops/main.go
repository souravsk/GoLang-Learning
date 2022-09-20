package main

import (
	"fmt"
)

func main(){
	fmt.Println("Wellcome to Golang loops\n")
	
	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Friday",
		"Saturday"}

	fmt.Println(days)

		//simple for loops

	// for i :=0; i < len(days); i++{  
	// 	fmt.Println(days[i])
	// }

		//for loops with range

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

		//
	for index, day := range days{
		fmt.Printf("Index is %v and value is %v\n",index, day)
	}
}