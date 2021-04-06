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

type CategoriesList []string

type ChuckFact struct {
	Categories []string
	Created    string `json:"created_at"`
	IconUrl    string `json:"icon_url"`
	Id         string
	Updated    string `json:"updated_at"`
	Url        string
	Value      string
}

var EmergencyFacts []string = []string{
	"Chuck Norris counted to infinity. Twice.",
	"Chuck Norris can strangle you with a cordless phone.",
	"Chuck Norris once went to mars. Thats why there are no signs of life.",
	"Chuck Norris is the reason Waldo is hiding.",
	"Chuck Norris will never have a heart attack... even a heart isnt foolish enough to attack Chuck Norris.",
}

type chuckFactList struct {
	Total  int
	Result []ChuckFact
}

const factURL string = "https://api.chucknorris.io/"

func RandomFact() (ChuckFact, error) {
	var fact ChuckFact
	responce, err := getAPI(factURL + "/jokes/random")
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

func Categorieslist() (CategoriesList, error) {
	responceList, err := getAPI(factURL + "jokes/categories")
	if err != nil {
		return nil, err
	}
	var categories CategoriesList
	err = json.Unmarshal(responceList, &categories)
	return categories, err
}

func RandomFactByCategory(category string) (ChuckFact, error) {
	var fact ChuckFact
	responce, err := getAPI(factURL + "/jokes/random?category=" + category)
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(responce, &fact)
	return fact, err
}

func RandomFactbytext(searhTerm string) (ChuckFact, error) {
	var factList chuckFactList
	var fact ChuckFact

	if strings.Contains(searhTerm, " ") {
		return fact, fmt.Errorf("search term must be one word")
	}

	responce, err := getAPI(factURL + "jokes/search?query=" + searhTerm)
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
