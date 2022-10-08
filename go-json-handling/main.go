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

	EncodeJson()
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
