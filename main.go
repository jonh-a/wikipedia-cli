package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

func formatSearchTerm(search string) string {
	return strings.Replace(search, " ", "_", -1)
}

func main() {
	searchFlag := flag.String("s", "", "term to search for")
	fullArticleFlag := flag.Bool("f", false, "show full article content")
	flag.Parse()

	searchValue := *searchFlag
	showFullArticle := *fullArticleFlag

	article, err := getSummary(formatSearchTerm(searchValue))
	if err.Error {
		log.Fatalln(err.Detail)
		os.Exit(1)
	}

	if showFullArticle {
		fullArticle, err := getFullArticle(article.Pageid)

		if err.Error {
			log.Fatalln(err.Detail)
			os.Exit(1)
		}

		out, renderErr := glamour.Render(fullArticle, "dark")

		if renderErr != nil {
			log.Fatal(renderErr)
		}

		fmt.Print(out)
		os.Exit(0)
	}

	fmt.Print(article.Extract)
	os.Exit(0)
}
