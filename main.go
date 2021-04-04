package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	responce, err := getJokesbytext("potato")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		fmt.Println(string(responce))
	}
}

const FactURL string = "https://api.chucknorris.io/"

func getRandomFact() ([]byte, error) {
	resp, err := http.Get(FactURL + "/jokes/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getCategorieslist() ([]byte, error) {
	resp, err := http.Get(FactURL + "jokes/categories")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getRandomFactByCategory(category string) ([]byte, error) {
	resp, err := http.Get(FactURL + "/jokes/random?category=" + category)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getJokesbytext(text string) ([]byte, error) {
	if strings.Contains(text, " ") {
		return nil, fmt.Errorf("text must be one word")
	}

	resp, err := http.Get(FactURL + "jokes/search?query=" + text)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
