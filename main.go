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

func getJSONDefinitions(input string, api dictionary.Dictionary) (string, error) {
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
		httpAddr = flag.String("http", "0.0.0.0:80", "HTTP service address.")
	)
	flag.Parse()

	api := dictionary.NewPearsonDictionaryDefault()

	// Use a buffered error channel so that handlers can
	// keep processing after throwing errors.
	errChan := make(chan error, 10)
	go func() {
		http.HandleFunc("/api/define", func (w http.ResponseWriter, r *http.Request) {
			var output string
			buf, err := ioutil.ReadAll(r.Body)

			if err == nil {
				output, err = getJSONDefinitions(string(buf), api)
			}

			if err != nil {
				fmt.Fprintf(w, "%q", html.EscapeString(err.Error()))
			} else {
				fmt.Fprintf(w, "%s", output)
			}
		})
		http.Handle("/", http.FileServer(http.Dir("./static")))

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
}