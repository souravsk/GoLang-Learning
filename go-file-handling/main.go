package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"io"
)

func main(){
	fmt.Println("Welcome to GoLang File Handling")

	context := "This is the text file that is created by using the go lang programming language Hope we all love it."

	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, context)
	checkNilErr(err)

	fmt.Println("Length is: ",length)
	defer file.Close()

	readfile("./myfile.txt")
}

func readfile(file string){
	datatype, err := ioutil.ReadFile(file)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n",string(datatype))
}

func checkNilErr (err error){
	if err != nil{
		panic(err)
	}
}