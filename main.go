package main

import (
	"fmt"

	"github.com/Anthonyntilelli/goChuckNorris/chuckApi"
)

func main() {
	fact, err := chuckApi.RandomFact()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		fmt.Println(fact)
	}
}
