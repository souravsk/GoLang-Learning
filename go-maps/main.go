package main

import (
	"fmt"
)
func main(){

// PROGRAM ON MAPS

	lang := make(map[string]string)

	lang["js"] = "javascript"
	lang["py"] = "python"
	lang["go"] = "Golang"
	lang["ts"] = "typescript"

	fmt.Println(lang)

	//lang = append(lang, ["MB"]"mongoDB")

	delete(lang, "ts")
	fmt.Println(lang)

	for keys, value := range lang {
		fmt.Printf("The is %v and the Value is %v\n", keys, value)
	}
}