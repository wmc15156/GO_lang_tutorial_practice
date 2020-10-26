package main

import (
	"fmt"

	"github.com/wmc15156/learngo/day04/mydict"
)

func main() {

	dict := mydict.Dictionary{}

	err := dict.Add("123", "greeting")

	if err != nil {
		fmt.Println(err)
	}

	data, err2 := dict.Search("123")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(data)
	}
}
