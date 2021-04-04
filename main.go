package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	responce, err := getRandomFact()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		fmt.Println(string(responce))
	}
}

const FactURL string = "https://api.chucknorris.io/"

func getRandomFact() ([]byte, error) {
	return getAPI(FactURL + "/jokes/random")
}

func getCategorieslist() ([]byte, error) {
	return getAPI(FactURL + "jokes/categories")
}

func getRandomFactByCategory(category string) ([]byte, error) {
	return getAPI(FactURL + "/jokes/random?category=" + category)
}

func getJokesbytext(text string) ([]byte, error) {
	if strings.Contains(text, " ") {
		return nil, fmt.Errorf("text must be one word")
	}

	return getAPI(FactURL + "jokes/search?query=" + text)
}

func getAPI(fullUrl string) ([]byte, error) {
	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	if (err == nil) && (resp.StatusCode != 200) {
		return nil, fmt.Errorf("Return status code is" + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
