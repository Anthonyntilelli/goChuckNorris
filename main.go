package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Anthonyntilelli/goChuckNorris/chuckApi"
)

func main() {
	// fact, err := chuckApi.RandomFact()
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fact)
	// }
	var selection string

	searchTerm := flag.String("s", "", "Get a fact based on search Term [choose only one flag]")
	category := flag.String("k", "", "Get a fact based on category [choose only one flag]")
	showCategories := flag.Bool("c", false, "Shows available categories [choose only one flag]")
	random := flag.Bool("r", false, "Shows a random fact [choose only one flag]")
	flag.Parse() // searchTerm, category, showCategories, random variables are not set till after parse

	if *searchTerm != "" {
		selection += "s"
	}
	if *category != "" {
		selection += "k"
	}
	if *showCategories {
		selection += "c"
	}
	if *random {
		selection += "r"
	}

	switch selection {
	case "s":
		fact, err := chuckApi.RandomFactbytext(*searchTerm)
		outputFact(fact, err)
	case "k":
		fact, err := chuckApi.RandomFactByCategory(*category)
		outputFact(fact, err)
	case "c":
		categories, err := chuckApi.CategoriesList()
		if err != nil {
			fmt.Println("ERROR: ", err)
			os.Exit(3)
		}
		println("Categories availble:")
		for _, cat := range categories {
			fmt.Println("  - ", cat)
		}
	case "r":
		fact, err := chuckApi.RandomFact()
		outputFact(fact, err)
	default:
		if len(selection) == 0 {
			fmt.Fprintln(os.Stderr, "Error: A flag must be selected")
		} else {
			fmt.Fprintln(os.Stderr, "Error: Please select only 1 flag")
		}
		flag.Usage()
		os.Exit(2)
	}
}

func outputFact(fact chuckApi.ChuckFact, err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("Emergeny Fact: ", chuckApi.EmergencyFact().Value)
		os.Exit(3)
	} else {
		fmt.Println(fact.Value)
	}
}
