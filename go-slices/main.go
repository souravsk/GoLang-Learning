package main

import(
	"fmt"
)
func main(){

	fmt.Println("Welcome to GO land Slices")
	var list = []string{"keyboard", "mouse", "moniter", "CPU"}

	fmt.Printf("The list is of %T\n", list)
	fmt.Println(append(list[:]))
	list = append(list, "cable", "powersuply")
	fmt.Println(list[:4])
	fmt.Println(len(list))
	fmt.Println(cap(list))

	datascores := make([]int, 6)

	datascores[0] = 234
	datascores[1] = 634
	datascores[2] = 434
	datascores[3] = 534
	datascores[4] = 734

	fmt.Println(datascores)
	datascores = append(datascores, 3444, 456, 4545)
	fmt.Println(datascores)
	fmt.Println(datascores[:2])
	fmt.Println(datascores[2:])

	sort.Ints(datascores)
	fmt.Println("Sorted data", datascores)
	//how to remove a value slices based on index

	var teckstack = []string{"c", "c++", "JAVA", "golang", "js"}
	fmt.Println(teckstack)

	var index int = 2                                             //we are using the index to delete the from the slices
	teckstack = append(teckstack[:index], teckstack[index+1:]...) // about this ... next time
	fmt.Println(teckstack)
}