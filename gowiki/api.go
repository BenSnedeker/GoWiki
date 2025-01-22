package gowiki

import "fmt"

// ------------
//    SEARCH
// ------------

// Params: Search query as string
// Return: A list of search results as strings presented by Wikipedia Search API
func WikiSearch(query string) ([]string, error) {
	// Get raw search response
	resp, err := fetchSearchResponse(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch search response: %v", err)
	}

	// Parse through search response and return results
	results, err := parseSearchResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse search results: %v", err)
	}

	return results, nil
}

func fetchSearchResponse(query string) (string, error) {
	// Build API URL

	// Fetch response

	return "", nil
}

func parseSearchResponse(resp string) ([]string, error) {
	// Split response into individual results

	// Clean each result

	return nil, nil
}

// ------------
//     READ
// ------------

// Params: Wikipedia article name as string
// Return: A string of the customized return text content
func WikiRead(articleName string) (string, error) {

	return "", nil
}
