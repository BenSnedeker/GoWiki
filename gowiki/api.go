package gowiki

import "fmt"

// ------------
//    SEARCH
// ------------

// Params: The search query as string
// Return: A list of search results as strings presented by Wikipedia Search API
func WikiSearch(flags Flags) ([]string, error) {
	// Fetch raw search response
	response, err := fetchSearchResponse(flags.Name)
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

//
func fetchSearchResponse(query string) (string, error) {
	// Build API URL

	// Fetch response

	return "", nil
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
func WikiRead(flags Flags) (string, error) {
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

func fetchArticleContent(flags Flags) (string, error) {
	// Build API URL

	// Fetch response

	return "", nil
}

func parseArticleContent(response string, flags Flags) (string, error) {
	// Check if no article page is found

	// Check if article suggestion page is found

	// Clean article

	return "", nil
}
