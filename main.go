package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func formatSearchTerm(search string) string {
	return strings.Replace(search, " ", "_", -1)
}

func sendRequest(search string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", search)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var j map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func main() {
	searchFlag := flag.String("s", "", "term to search for")
	flag.Parse()

	searchValue := *searchFlag

	data, err := sendRequest(formatSearchTerm(searchValue))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data["extract"])
}
