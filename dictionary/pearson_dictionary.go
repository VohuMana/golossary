package dictionary

import (
	"errors"
	"net/url"
	"net/http"
	"encoding/json"
	"fmt"
)

// PearsonDictionary implements the Dictionary interface and is used to call the Pearson Dictionary API's
type PearsonDictionary struct {
	endpoint string
	activeDictionary string
	cachedResults []PearsonDefineWordContract
}

// NewPearsonDictionaryDefault creates a new dictionary using the Pearson api and the Longman Dictionary of Contemporary English (5th edition)
func NewPearsonDictionaryDefault() *PearsonDictionary {
	return &PearsonDictionary {
		endpoint: "http://api.pearson.com/v2/dictionaries",
		activeDictionary: "ldoce5",
	}
}

// NewPearsonDictionaryCustom creates a new dictionary using the provided endpoint and dictionary code while conforming to the Pearson API
func NewPearsonDictionaryCustom(baseEndpoint, dictionaryCode string) *PearsonDictionary {
	return &PearsonDictionary {
		endpoint: baseEndpoint,
		activeDictionary: dictionaryCode,
	}
}

// DefineWord gets all definitions of a word from the Pearson Dictionary API
func (p *PearsonDictionary) DefineWord(word string) ([]string, error) {
	var err error
	var response *http.Response
	var parsedResult PearsonDefineWordContract
	var definitions []string

	// Call the API
	pearsonURL, err := p.constructDefineWordURL(word)

	if err == nil {
		response, err = http.Get(pearsonURL)
	}

	if err == nil {
		if response.Body != nil {
			defer response.Body.Close()
		}

		if response.StatusCode != 200 {
			err = fmt.Errorf("Request failed with HTTP: %v", response.StatusCode)
		}
	}

	// Parse the result
	if err == nil {
		err = json.NewDecoder(response.Body).Decode(&parsedResult)
	}

	// Save the data
	if err == nil {
		p.cachedResults = append(p.cachedResults, parsedResult)
	}

	// Ensure results were returned
	if err == nil {
		if parsedResult.Count == 0 {
			err = errors.New("No results for given word")
		}
	}

	// Search for the word that matches the passed in word
	if err == nil {
		for _, result := range parsedResult.Results {
			if result.Headword == word {
				for _, sense := range result.Senses {
					definitions = sense.Definition
					break
				}
				break
			}
		}
	}
	// Return that result
	return definitions, err
}

func (p *PearsonDictionary) constructDefineWordURL(word string) (string, error) {
	var err error
	URL := fmt.Sprintf("%s/%s/entries", p.endpoint, p.activeDictionary)
	baseURL, err := url.Parse(URL)

	if err == nil {
		queryParams := url.Values{}
		queryParams.Add("headword", word)
		baseURL.RawQuery = queryParams.Encode()
	}

	return baseURL.String(), err
}