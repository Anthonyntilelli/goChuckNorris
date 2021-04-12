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
		fmt.Println("search Term")
	case "k":
		fmt.Println("Category Search")
	case "c":
		fmt.Println("List categories")
	case "r":
		fmt.Println("random")
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
