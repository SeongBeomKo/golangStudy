package main

import (
	"Dictionary/myDict"
	"fmt"
)

func main() {
	dictionary := myDict.Dictionary{"first": "first word"}
	dictionary["hello"] = "hello"

	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	fmt.Println(dictionary)
}
