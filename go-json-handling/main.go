package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Coursename"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Welcome To Go Lang Json Handling")

	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"HTML BootCamp", 1000, "youtube.com/freecodecamp", "thikhai", []string{"web-dev", "basic"}},
		{"CSS BootCamp", 10000, "youtube.com/freecodecamp", "achahai", []string{"style", "basic"}},
		{"JS BootCamp", 100000, "youtube.com/freecodecamp", "kuchthiknahihai", []string{"backend", "basic"}},
	}

	//package this data as  json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonData := []byte(`

		{
			"Coursename": "HTML BootCamp",
			"Price": 1000,
			"Website": "youtube.com/freecodecamp",
			"tags": ["web-dev","basic"]
		}
	`)

	var lcoCourses course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("NOT A Valid JSON")
	}

	//same case where you just want to add data to key value pair

	var mydata map[string]interface{}
	//we are using map because we need key value pair the string for key because it alwaye a string bur we don't know what kind of data can be in a value so we use interface method
	json.Unmarshal(jsonData, &mydata)
	fmt.Printf("%#v\n", mydata)

}
