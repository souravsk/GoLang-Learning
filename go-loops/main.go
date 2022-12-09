package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Wellcome to Golang loops")
	
	/* days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Friday",
		"Saturday"}

	fmt.Println(days) */

		//simple for loops

	// for i :=0; i < len(days); i++{  
	// 	fmt.Println(days[i])
	// }

		//for loops with range

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

		//
	/* for index, day := range days{
		fmt.Printf("Index is %v and value is %v\n",index, day)
	} */

	var randam = rand.Intn(101)

	if randam == 50 {
		fmt.Println("it 50")
	}
	
	if randam > 50{
		fmt.Println("It's closer to 100",randam)
	} else if randam < 50{
		fmt.Println("It's closer to 0",randam)
	} else{
		fmt.Println(randam)	
	}

}