// Package chuckApi makes https calls to the Chuck Norris Api to get a facts.
// Each call is done once, retries need be handles outside the package.
package chuckApi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ChuckFact contains the fact with some metadata convered from API Json reponse.
// If more metadate is required, uncomment the extra entries.
// However, `Categories`, `Id`, and `Value` are required.
type ChuckFact struct {
	Categories []string
	Id         string
	Value      string
	// Created    string `json:"created_at"`
	// IconUrl    string `json:"icon_url"`
	// Updated    string `json:"updated_at"`
	// Url        string
}

type chuckFactList struct {
	Total  int
	Result []ChuckFact
}

// factURL points to api end point.
// Must start with `https` and end with a `/`
const factURL string = "https://api.chucknorris.io/"

func RandomFact() (ChuckFact, error) {
	var fact ChuckFact
	responce, _, err := getAPI(factURL + "/jokes/random")
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

func CategoriesList() ([]string, error) {
	responceList, _, err := getAPI(factURL + "jokes/categories")
	if err != nil {
		return nil, err
	}
	var categories []string
	err = json.Unmarshal(responceList, &categories)
	return categories, err
}

// RandomFactByCategory does not check locally, relies on api for errors
func RandomFactByCategory(category string) (ChuckFact, error) {
	var fact ChuckFact
	responce, StatusCode, err := getAPI(factURL + "/jokes/random?category=" + category)
	if err != nil {
		if StatusCode == 404 {
			err = fmt.Errorf("Provided category (" + category + ") is not valid.")
		}
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

// RandomFactbytext looks for facts based on search term and returns one randomly from the list.
// searchTerm MUST be one word
// If no facts can be found for a given search term, an error is returned.
func RandomFactbytext(searhTerm string) (ChuckFact, error) {
	var factList chuckFactList
	var fact ChuckFact

	if strings.Contains(searhTerm, " ") {
		return fact, fmt.Errorf("search term must be one word")
	}

	responce, _, err := getAPI(factURL + "jokes/search?query=" + searhTerm)
	if err != nil {
		return fact, err
	}
	err = json.Unmarshal(responce, &factList)
	if err != nil {
		return fact, err
	}
	if len(factList.Result) == 0 {
		return fact, fmt.Errorf("no results for provided searhTerm")
	}
	// Pick one fact from list randomly
	min, max := 0, len(factList.Result)-1
	pick := min + rand.Intn(max-min)
	fact = factList.Result[pick]
	return fact, err
}

// EmergencyFact will return a valid ChuckFact from a small local list in the function.
func EmergencyFact() ChuckFact {
	emergencyFacts := [...]string{
		"Chuck Norris counted to infinity. Twice.",
		"Chuck Norris can strangle you with a cordless phone.",
		"Chuck Norris once went to mars. Thats why there are no signs of life.",
		"Chuck Norris is the reason Waldo is hiding.",
		"Chuck Norris will never have a heart attack... even a heart isnt foolish enough to attack Chuck Norris.",
	}

	// Pick one fact from list randomly
	min, max := 0, len(emergencyFacts)-1
	pick := min + rand.Intn(max-min)
	value := emergencyFacts[pick]

	return ChuckFact{
		Categories: nil,
		Id:         "{Emergency}",
		Value:      value,
	}
}

// Valid only check ID and Value.
func (c ChuckFact) Valid() bool {
	return c.Id != "" && c.Value != ""
}

func getAPI(fullUrl string) ([]byte, int, error) {
	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, 0, err
	}
	if (err == nil) && (resp.StatusCode != http.StatusOK) {
		return nil, resp.StatusCode, fmt.Errorf("Return status code is: " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
