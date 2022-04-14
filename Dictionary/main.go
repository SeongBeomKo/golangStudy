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

	err1 := dictionary.Add("hi", "Greeting")
	if err1 != nil {
		fmt.Println(err1)
	}

	err3 := dictionary.Update("hi", "greeting")
	if err3 != nil {
		fmt.Println(err3)
	}

	hi, err2 := dictionary.Search("hi")

	if err2 == nil {
		fmt.Println(hi)
	}

	dictionary.Delete("hello")
	word, err4 := dictionary.Search("hello")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(word)
	}

}
