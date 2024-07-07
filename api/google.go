package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Kind       string `json:"kind"`
	Items      []Item `json:"items"`
	SearchInfo struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
}

type Item struct {
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	HTMLTitle   string `json:"htmlTitle"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
	HTMLSnippet string `json:"htmlSnippet"`
}

type ReturnData struct {
	Title string
	Link string
}

func GetGoogleResponse(query string) ([]ReturnData, error){
	er := godotenv.Load()
	if er != nil {
		return nil, errors.New("error loading .env file")
	}

	GOOGLE_API_KEY := os.Getenv("GOOGLE_API_KEY")
	SEARCH_ENGINE_ID := os.Getenv("SEARCH_ENGINE_ID")

	searchURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", GOOGLE_API_KEY, SEARCH_ENGINE_ID, url.QueryEscape(query))

	var response *http.Response
	var err error

	// Retry mechanism
	for i := 0; i < 3; i++ {
		response, err = http.Get(searchURL)
		if err == nil {
			break
		}
		fmt.Println("Error making the request, retrying:", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		fmt.Println("Error making the request after retries:", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling the response body:", err)
		return nil, err
	}

	var Results []ReturnData

	for _, item := range result.Items {
		rd := ReturnData{
			Title: item.Title,
			Link:  item.Link,
		}
		Results = append(Results, rd)
	}
	return Results, nil
}
