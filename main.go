package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	responce, _ := getRandomFact(FactURL)
	fmt.Println(string(responce))
}

const FactURL string = "https://api.chucknorris.io/"

func getRandomFact(url string) ([]byte, error) {
	resp, err := http.Get(url + "/jokes/random")
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
