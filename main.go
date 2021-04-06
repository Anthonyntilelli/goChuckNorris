package main

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

func main() {
	fact, err := getRandomFactbytext("potato")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		fmt.Println(fact)
	}
}

const FactURL string = "https://api.chucknorris.io/"

type ChuckFact struct {
	Categories []string
	Created    string `json:"created_at"`
	IconUrl    string `json:"icon_url"`
	Id         string
	Updated    string `json:"updated_at"`
	Url        string
	Value      string
}

type chuckFactList struct {
	Total  int
	Result []ChuckFact
}

type categoriesList []string

var EmergencyFacts []string = []string{
	"Chuck Norris counted to infinity. Twice.",
	"Chuck Norris can strangle you with a cordless phone.",
	"Chuck Norris once went to mars. Thats why there are no signs of life.",
	"Chuck Norris is the reason Waldo is hiding.",
	"Chuck Norris will never have a heart attack... even a heart isnt foolish enough to attack Chuck Norris.",
}

func getRandomFact() (ChuckFact, error) {
	var fact ChuckFact
	responce, err := getAPI(FactURL + "/jokes/random")
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

func getCategorieslist() (categoriesList, error) {
	responceList, err := getAPI(FactURL + "jokes/categories")
	if err != nil {
		return nil, err
	}
	var categories categoriesList
	err = json.Unmarshal(responceList, &categories)
	return categories, err
}

func getRandomFactByCategory(category string) (ChuckFact, error) {
	var fact ChuckFact
	responce, err := getAPI(FactURL + "/jokes/random?category=" + category)
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

func getRandomFactbytext(searhTerm string) (ChuckFact, error) {
	var factList chuckFactList
	var fact ChuckFact

	if strings.Contains(searhTerm, " ") {
		return fact, fmt.Errorf("search term must be one word")
	}

	responce, err := getAPI(FactURL + "jokes/search?query=" + searhTerm)
	if err != nil {
		return fact, err
	}
	err = json.Unmarshal(responce, &factList)
	if err != nil {
		return fact, err
	}
	// Pick one fact from list randomly
	rand.Seed(time.Now().UnixNano())
	min, max := 0, len(factList.Result)-1
	pick := min + rand.Intn(max-min)
	fact = factList.Result[pick]
	return fact, err
}

func getAPI(fullUrl string) ([]byte, error) {
	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	if (err == nil) && (resp.StatusCode != http.StatusOK) {
		return nil, fmt.Errorf("Return status code is: " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
