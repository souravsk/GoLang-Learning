package main

import (
	"fmt"
	"sort"
)

func main() {
	// welcome := "welcome to user input"
	// fmt.Printf(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our Pizza:-")

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for the reating", input)

	// PROGRAM FOR SI

	// fmt.Print("Enter the Pricpal Amount;- ")
	// var p int
	// fmt.Scanf("%d", &p)
	// fmt.Print("Enter the Loan Intrated:- ")
	// var i int
	// fmt.Scanf("%d", &i)
	// fmt.Print("Enter the Loan Tenur :- ")
	// var t int
	// fmt.Scanf("%d", &t)
	// simpleIntret := (p * t * i) / 100
	// fmt.Println("Simple Intrest is :-", simpleIntret)

	//PROGRAM FOR CALCULATING

	// fmt.Println("Enter any two number to form a calculation")
	// fmt.Print("Enter the first Number:- ")
	// var a int
	// fmt.Scanf("%d", &a)
	// fmt.Print("Enter tht Seconad Number:- ")
	// var b int
	// fmt.Scanf("%d", &b)
	// fmt.Print("Select any Operation + , - , / , *  :-  ")
	// var op string
	// fmt.Scanf("%s", &op)
	// var result int
	// if op == "+" {
	// 	result = a + b
	// 	fmt.Printf("Add of the number is %d \n", result)
	// } else if op == "-" {
	// 	result = a - b
	// 	fmt.Printf("Sub of the Number is %d \n", result)
	// } else if op == "/" {
	// 	result = a / b
	// 	fmt.Printf("Div of the Number is %d \n", result)
	// } else if op == "*" {
	// 	result = a * b
	// 	fmt.Printf("Muiti of the Number is %d \n", result)
	// } else {
	// 	fmt.Println("Please Enter the right Operater ")
	// }

	// PROGRAM FOR LARGEST NUMBER

	// fmt.Println("Comparing Two Numbers")
	// fmt.Print("Enter Your First Number:- ")
	// var num1 int
	// fmt.Scan(&num1)
	// fmt.Print("Enter Your Seconad Number:- ")
	// var num2 int
	// fmt.Scan(&num2)
	// if num1 > num2 {
	// 	fmt.Printf("First Number is Bigger %d \n", num1)
	// } else if num2 > num1 {
	// 	fmt.Printf("Seconad number is Bigger %d \n", num2)
	// } else {
	// 	fmt.Printf("Both Of the Numbers are Equle %d and %d \n", num1, num2)
	// }

	// PROGRAM TO CONVERT THE INR TO USD

	// fmt.Println("Converting the INR TO USD")
	// fmt.Print("Enter the Amount in INR:- ")
	// var INR float32
	// fmt.Scan(&INR)
	// var USD float32
	// USD = INR * 80
	// fmt.Printf("Your Amount in USD is %v \n", USD)

	// PROGRAM TO GET THE TIME

	// fmt.Println("Program about Time")
	// today := time.Now()
	// fmt.Println("Today's Date, Time, and Location is ", today)

	// fmt.Println(today.Format("01-02-2006 Monday 15:04:05"))

	// PROGRAM ON POINTERS
	// var pointer *int
	// var variable int = 1000
	// fmt.Println("Just create without any valuse and just variable name", pointer)
	// // fmt.Println("Just create without any valuse and with * ", *pointer)
	// // fmt.Println("Just create without any valuse and with & ", &pointer)

	// pointer = &variable
	// fmt.Printf("it may show some garbage value %v \n", pointer)
	// fmt.Printf("it will show the value in the variable * %v \n", *pointer)
	// fmt.Printf("it will show the address of variable  & %v \n", &pointer)

	//POGRAM ON SLICES
	// var list = []string{"keyboard", "mouse", "moniter", "CPU"}

	// fmt.Printf("The list is of %T\n", list)
	// fmt.Println(append(list[:]))
	// list = append(list, "cable", "powersuply")
	// fmt.Println(list[:4])
	// fmt.Println(len(list))
	// fmt.Println(cap(list))

	// datascores := make([]int, 6)

	// datascores[0] = 234
	// datascores[1] = 634
	// datascores[2] = 434
	// datascores[3] = 534
	// datascores[4] = 734

	// fmt.Println(datascores)
	// datascores = append(datascores, 3444, 456, 4545)
	// fmt.Println(datascores)
	// fmt.Println(datascores[:2])
	// fmt.Println(datascores[2:])

	// sort.Ints(datascores)
	// fmt.Println("Sorted data", datascores)
	// //how to remove a value slices based on index

	// var teckstack = []string{"c", "c++", "JAVA", "golang", "js"}
	// fmt.Println(teckstack)

	// var index int = 2                                             //we are using the index to delete the from the slices
	// teckstack = append(teckstack[:index], teckstack[index+1:]...) // about this ... next time
	// fmt.Println(teckstack)

	// // PROGRAM ON MAPS

	// lang := make(map[string]string)

	// lang["js"] = "javascript"
	// lang["py"] = "python"
	// lang["go"] = "Golang"
	// lang["ts"] = "typescript"

	// fmt.Println(lang)

	// //lang = append(lang, ["MB"]"mongoDB")

	// delete(lang, "ts")
	// fmt.Println(lang)

	// for keys, value := range lang {
	// 	fmt.Printf("The is %v and the Value is %v\n", keys, value)
	// }

	// STRUCTS  IN GOLANG

	sonu := User("sonu", "sonuak47@gmail.com", true, 14)
	fmt.Print(sonu)

}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}