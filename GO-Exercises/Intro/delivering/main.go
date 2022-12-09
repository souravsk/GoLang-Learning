package main

import (
	"fmt"
	"reflect"
)

/* func main(){
	var numOfPackage = 100

	fmt.Printf("I have %v of Package \n",numOfPackage)
	numOfPackage -= 20
	fmt.Printf("After delivering 20 Pacakges i have to now %v \n",numOfPackage)
	numOfPackage /= 4
	fmt.Printf("Packages are divided into 4 customers of each customer got %v packages \n", numOfPackage)
} */

func main() {
    var food = "Pizza"
    var slices = 4
    var pineappleOnPizza = false

	fmt.Printf("The Food Variable is of %v type \n",reflect.TypeOf(food))
	fmt.Printf("The Slices Variable is of %v type \n",reflect.TypeOf(slices))
	fmt.Printf("The pineappleOnPizza Variable is of %v type \n",reflect.TypeOf(pineappleOnPizza))
}