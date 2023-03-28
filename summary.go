package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Pageid  int64  `json:"pageid"`
	Extract string `json:"extract"`
}

type SearchError struct {
	Detail     string `json:"detail"`
	Uri        string `json:"uri"`
	Error      bool
	ErrorValue error
}

func getSummary(search string) (Article, SearchError) {
	url := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", search)

	resp, err := http.Get(url)
	if err != nil {
		return Article{}, SearchError{Error: true, ErrorValue: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var a Article
		err = json.NewDecoder(resp.Body).Decode(&a)
		if err != nil {
			fmt.Println(err)
			return Article{}, SearchError{Error: true, ErrorValue: err}
		}

		return a, SearchError{}
	}

	if resp.StatusCode == 404 {
		var se SearchError
		err = json.NewDecoder(resp.Body).Decode(&se)

		if err != nil {
			fmt.Println(err)
			return Article{}, SearchError{Error: true, ErrorValue: err}
		}

		se.Error = true
		return Article{}, se
	}

	return Article{}, SearchError{}
}
