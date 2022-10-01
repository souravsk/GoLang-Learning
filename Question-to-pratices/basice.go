package main

import (
	"fmt"
)

func main(){
	radius := 10.0
	
	area := 2*radius*3.14

	fmt.Println("The area of the circle - ",area)

	arraysum()
}

func arraysum(){
	list := make([]int, 0, 5)
	list = append(list, 5, 6, 8, 9, 12)

	data := []int{45,67,4,6,3}
	fmt.Println(data)

	datasize := make([]int, 5)
	datacopyed := copy(datasize,data)
	fmt.Println(datacopyed)

	sum := 0

	fmt.Println(list)

	for i := range list{
		sum += list[i]
	}

	fmt.Println("The sum of slice is ", sum)

}