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
	if searchTerm != "" {
		set++
		fmt.Println("set ", set) //DIAG
	}
	flag.StringVar(&category, "k", "", "Get a fact based on category [choose only one flag]")
	if category != "" {
		set++
		fmt.Println("set ", set) //DIAG
	}
	flag.BoolVar(&showCategories, "c", false, "Shows available categories [choose only one flag]")
	if showCategories {
		set++
		fmt.Println("set ", set) //DIAG
	}
	flag.BoolVar(&random, "r", false, "Shows a random fact [choose only one flag]")
	if random {
		set++
		fmt.Println("set ", set) //DIAG
	}
	flag.Parse()

	//DIAG
	fmt.Print("searchTerm ", searchTerm, "\n")
	fmt.Print("category ", category, "\n")
	fmt.Print("ShowCategories ", showCategories, "\n")
	fmt.Print("random ", random, "\n")
	fmt.Print("set ", set, "\n")
	//DIAG

	if set == 0 {
		flag.Usage()
		os.Exit(2)
	}
	if set != 1 {
		fmt.Fprintf(os.Stderr, "Only Select 1 flag\n")
	}
}
