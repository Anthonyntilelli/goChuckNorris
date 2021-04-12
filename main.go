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
		if err != nil {
			// Replace error if catigory is invalid
			valid, vErr := validCategory(*category)
			if !valid && vErr == nil {
				err = fmt.Errorf(*category + " is not a valid category")
			}
		}
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
		emergenyExit(err)
	} else {
		fmt.Println(fact.Value)
	}
}

func emergenyExit(err error) {
	fmt.Fprintln(os.Stderr, "ERROR: ", err)
	fmt.Println("Emergeny Fact: " + chuckApi.EmergencyFact().Value)
	os.Exit(3)
}

// Use original error if validCategory returns error
func validCategory(input string) (bool, error) {
	// check for valid category
	list, err := chuckApi.CategoriesList()
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v == input {
			return true, nil
		}
	}
	return false, nil
}
