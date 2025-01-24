package gowiki

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ------------
//    SEARCH
// ------------

// Params: The search query as string
// Return: A list of search results as strings presented by Wikipedia Search API
func WikiSearch(flags *Flags) ([]string, error) {
	// Fetch raw search response
	response, err := fetchSearchResponse(flags)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch search response: %v", err)
	}

	// Parse through search response and return results
	results, err := parseSearchResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse search results: %v", err)
	}

	return results, nil
}

func fetchSearchResponse(flags *Flags) (string, error) {
	// Build API URL

	// Base URL to be used
	baseURL := "https://en.wikipedia.org/w/api.php"

	// Establish parameters
	params := url.Values{}
	params.Add("action", "query")
	params.Add("list", "search")
	params.Add("srsearch", flags.Term)
	params.Add("format", "json")

	// Build full URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Fetch response
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch HTML: %v", err)
	}
	defer resp.Body.Close()

	// Ensure status 200
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}

func parseSearchResponse(response string) ([]string, error) {
	// Split response into individual results

	// Clean each result

	return nil, nil
}

// ------------
//     READ
// ------------

// Params: The Wikipedia article name as string
// Return: A string of the customized text content
func WikiRead(flags *Flags) (string, error) {
	// Fetch raw API result for
	response, err := fetchArticleContent(flags)
	if err != nil {
		return "", fmt.Errorf("failed to fetch article content: %v", err)
	}

	// Parse through article content
	content, err := parseArticleContent(response, flags)
	if err != nil {
		return "", fmt.Errorf("failed to parse article content: %v", err)
	}

	return content, nil
}

func fetchArticleContent(flags *Flags) (string, error) {
	// Build API URL

	// Fetch response

	return "", nil
}

func parseArticleContent(response string, flags *Flags) (string, error) {
	// Check if no article page is found

	// Check if article suggestion page is found

	// Clean article

	return "", nil
}
