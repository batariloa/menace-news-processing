package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/batariloa/bobber/internal/datastruct"
)

var apiKey = `c0ab468f-f74f-4559-9427-9a82061a5b4c`

func GetRecentArticles() ([]datastruct.FetchedIncident, error) {
	requestBody := datastruct.NewsRequestBody{

		Action:                 "getArticles",
		Keyword:                "Florida man",
		ArticlesPage:           1,
		ArticlesCount:          10,
		ArticlesSortBy:         "date",
		ArticlesSortByAsc:      false,
		ArticlesArticleBodyLen: -1,
		ResultType:             "articles",
		DataType:               []string{"news", "pr"},
		ApiKey:                 apiKey,
		ForceMaxDataTimeWindow: 30,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	url := `http://eventregistry.org/api/v1/article/getArticles`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("FW: Error requesting articles.", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	fmt.Println("Raw JSON:", string(body))

	var fetchedResults datastruct.FetchedResults
	if err := json.NewDecoder(resp.Body).Decode(&fetchedResults); err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return nil, err
	}

	return fetchedResults.Articles.Results, nil
}
