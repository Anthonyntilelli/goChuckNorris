package main

import (
	"flag"
	"fmt"
	"os"
)

// "github.com/Anthonyntilelli/goChuckNorris/chuckApi"

func main() {
	// fact, err := chuckApi.RandomFact()
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fact)
	// }
	var searchTerm, category string
	var showCategories, random bool
	set := 0

	flag.StringVar(&searchTerm, "s", "", "Get a fact based on search Term [choose only one flag]")
	flag.StringVar(&category, "k", "", "Get a fact based on category [choose only one flag]")
	flag.BoolVar(&showCategories, "c", false, "Shows available categories [choose only one flag]")
	flag.BoolVar(&random, "r", false, "Shows a random fact [choose only one flag]")
	flag.Parse() // searchTerm, category, showCategories, random variables are not set till after parse

	if searchTerm != "" {
		set++
	}
	if category != "" {
		set++
	}
	if showCategories {
		set++
	}
	if random {
		set++
	}
	if set == 0 {
		fmt.Fprintln(os.Stderr, "Please select 1 flag")
		flag.Usage()
		os.Exit(2)
	}
	if set != 1 {
		fmt.Fprintln(os.Stderr, "Error: Only Select 1 flag.")
		os.Exit(2)
	}
}
