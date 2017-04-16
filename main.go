package main

import (
	"fmt"
	"github.com/vohumana/golossary/dictionary"
)

func main() {
	api := dictionary.NewPearsonDictionaryDefault()

	definitions, err := api.DefineWord("cool")
	if err != nil {
		panic(err)
	}

	fmt.Println(definitions)
}