// Issues prints a summary table of github issues based on the provided
// search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rigikulm/gopl/issues/internal/github"
)

func SearchIssues(terms []string) ([]string, error) {
	issues := []string{"This is issue 1", "This is issue 2"}
	return issues, nil
}

func main() {
	searchTerms := os.Args[1:]
	fmt.Println("Welcome to `issues`!")
	fmt.Printf("Search Terms: %v\n", searchTerms)

	result, err := github.SearchIssues(searchTerms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("Issue: %v\n", item)
	}

	log.Print("That's all folks!")
}