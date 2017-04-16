package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/vohumana/golossary/dictionary"
	"github.com/vohumana/golossary/tokenizer"
)

func main() {
	api := dictionary.NewPearsonDictionaryDefault()

	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	tok := tokenizer.NewEnglish(string(buf))

	words := tok.GetTokens()

	uniqueWords := make(map[string][]string)

	for _, word := range words {
		uniqueWords[strings.ToLower(word)] = nil
	}

	for word := range uniqueWords {
		fmt.Printf("Looking up definition for %s...\n", word)
		def, err := api.DefineWord(word)
		if err != nil {
			fmt.Println(err)
		} else {
			uniqueWords[word] = def
		}
	}
	
	fmt.Println("")

	wordsWithNoDefinition := make([]string, 1)

	for word, definitions := range uniqueWords {
		if len(definitions) != 0 {
			fmt.Printf("%s\n", word)
			for _, def := range definitions {
				fmt.Printf("---> %s\n", def)
			}
		} else {
			wordsWithNoDefinition = append(wordsWithNoDefinition, word)
		}
	}

	if len(wordsWithNoDefinition) != 0 {
		fmt.Println("\nCould not find definitions for these words:")

		for _, word := range wordsWithNoDefinition {
			fmt.Println(word)
		}
	}
}