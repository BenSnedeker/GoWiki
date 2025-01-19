package gowiki

import (
	"fmt"
	"io"
	"net/http"
)

// -------------
//    SEARCH
// -------------

func WikiSearch(query string) (string, error) {
	// Fetch HTML response
	htmlResp, err := getSearchResponse(query)
	if err != nil {
		return "", fmt.Errorf("failed to search Wikipedia: %v", err)
	}

	// Clean HTML into string
	searchResults, err := CleanSearchHTML(htmlResp)
	if err != nil {
		return "", fmt.Errorf("failed to clean search results: %v", err)
	}

	return searchResults, nil
}

func getSearchResponse(query string) (string, error) {
	// Build URL
	baseURL := "https://en.wikipedia.org/w/api.php"
	params := fmt.Sprintf("?action=query&list=search&srsearch=%s&format=json", query)
	fullURL := baseURL + params

	// Make HTTP GET request to Wikipedia
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch search results: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	// Get response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

// -------------
//     READ
// -------------

func WikiRead(flags Flags, query string) string {
	fmt.Printf("Opening article \033[1;34m\"%s\"\033[0m\n", query)

	return "NOT IMPLEMENTED"
}
