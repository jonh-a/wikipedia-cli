package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	if showFullArticle {
		fullArticle, err := getFullArticle(formatSearchTerm(searchValue))
		if err.Error {
			fmt.Println(err.Detail)
			os.Exit(1)
		}

		fmt.Println(fullArticle)
		os.Exit(0)
	}

	if !showFullArticle {
		article, err := getSummary(formatSearchTerm(searchValue))
		if err.Error {
			fmt.Println(err.Detail)
			os.Exit(1)
		}

		fmt.Println(article.Extract)
	}
}
