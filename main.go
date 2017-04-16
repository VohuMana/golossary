package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"flag"
	"net/http"
	"html"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vohumana/golossary/dictionary"
	"github.com/vohumana/golossary/tokenizer"
)

func getJSONDefinitions(input string) (string, error) {
	api := dictionary.NewPearsonDictionaryDefault()
	tok := tokenizer.NewEnglish(input)
	words := tok.GetTokens()
	uniqueWords := make(map[string][]string)

	for _, word := range words {
		uniqueWords[strings.ToLower(word)] = nil
	}

	for word := range uniqueWords {
		def, err := api.DefineWord(word)
		if err != nil {
			fmt.Println(err)
		} else {
			uniqueWords[word] = def
		}
	}

	jsonOut,err := json.Marshal(uniqueWords)

	return string(jsonOut), err
}

func main() {
	var (
		httpAddr = flag.String("http", ":80", "HTTP service address.")
	)
	flag.Parse()

	// Use a buffered error channel so that handlers can
	// keep processing after throwing errors.
	errChan := make(chan error, 10)
	go func() {
		http.HandleFunc("/api/define", func (w http.ResponseWriter, r *http.Request) {
			var output string
			buf, err := ioutil.ReadAll(r.Body)

			if err == nil {
				output, err = getJSONDefinitions(string(buf))
			}

			if err != nil {
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(err.Error()))
			} else {
				fmt.Fprintf(w, "%s", output)
			}
		})

		log.Println("Starting server...")
		log.Printf("HTTP service listening on %s", *httpAddr)

		errChan <- http.ListenAndServe(*httpAddr, nil)
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			// Log any errors from our server
			log.Fatal(err)
		case s := <-signalChan:
			// ctrl+c is a clean exit
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			os.Exit(0)
		}
}

	// api := dictionary.NewPearsonDictionaryDefault()
	
	// fileName := flag.String("f", "/usr/bin/test.txt", "File to parse")
	// flag.Parse()

	// buf, err := ioutil.ReadFile(*fileName)
	// if err != nil {
	// 	panic(err)
	// }

	// tok := tokenizer.NewEnglish(string(buf))

	// words := tok.GetTokens()

	// uniqueWords := make(map[string][]string)

	// for _, word := range words {
	// 	uniqueWords[strings.ToLower(word)] = nil
	// }

	// for word := range uniqueWords {
	// 	fmt.Printf("Looking up definition for %s...\n", word)
	// 	def, err := api.DefineWord(word)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		uniqueWords[word] = def
	// 	}
	// }
	
	// fmt.Println("")

	// wordsWithNoDefinition := make([]string, 1)

	// for word, definitions := range uniqueWords {
	// 	if len(definitions) != 0 {
	// 		fmt.Printf("%s\n", word)
	// 		for _, def := range definitions {
	// 			fmt.Printf("---> %s\n", def)
	// 		}
	// 	} else {
	// 		wordsWithNoDefinition = append(wordsWithNoDefinition, word)
	// 	}
	// }

	// if len(wordsWithNoDefinition) != 0 {
	// 	fmt.Println("\nCould not find definitions for these words:")

	// 	for _, word := range wordsWithNoDefinition {
	// 		fmt.Println(word)
	// 	}
	// }
}