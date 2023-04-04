package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/microcosm-cc/bluemonday"
)

func convertToMarkdown(input string) string {
	policy := bluemonday.UGCPolicy()
	html := policy.Sanitize(input)
	converter := md.NewConverter("", true, nil)

	md, err := converter.ConvertString(html)

	if err != nil {
		log.Fatal(err)
	}

	return string(md)
}

func getExtract(a map[string]interface{}) (string, error) {
	query, ok := a["query"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("missing query map")
	}

	pages, ok := query["pages"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("missing pages map")
	}

	var pageID string
	for k := range pages {
		pageID = k
		break
	}

	page, ok := pages[pageID].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("missing page map")
	}

	extract, ok := page["extract"].(string)
	if !ok {
		return "", fmt.Errorf("missing extract value")
	}

	md := convertToMarkdown(extract)

	return md, nil
}

func getFullArticle(search string) (string, SearchError) {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=%s&exsectionformat=wiki", search)

	resp, err := http.Get(url)
	if err != nil {
		return "", SearchError{Error: true}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var r map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&r)

		if err != nil {
			fmt.Println(err)
			return "", SearchError{Error: true}
		}

		extract, _ := getExtract(r)

		return extract, SearchError{Error: false}
	}

	return "", SearchError{Error: false}
}
