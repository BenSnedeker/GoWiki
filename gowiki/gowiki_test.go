package gowiki

import (
	"fmt"
	"testing"
)

// -------------
//   WikiSearch
// -------------

// Mock func to replace `getSearchResponse`
func mockGetSearchResponse(query string) (string, error) {
	if query == "valid query" {
		return `<html><body><div class="result">Search Result</div></body></html>`, nil
	} else if query == "empty query" {
		return "", fmt.Errorf("emptry query provided")
	}
	return "", fmt.Errorf("unexpected error")
}

// Mock fun to replace `CleanSearchHTML`
func mockCleanSearchHTML(html string) (string, error) {
	if html == `<html><body><div class="result">Search Result</div></body></html>` {
		return "Search Result", nil
	}
	return "", fmt.Errorf("failed to clean HTML")
}

// Test `WikiSearch`
func TestWikiSearch(t *testing.T) {
	// Backup original functions
	originalGetSearchResponse := getSearchResponse
	originalCleanSearchHTML := CleanSearchHTML

	// Defer restoration of originals
	defer func() {
		getSearchResponse = originalGetSearchResponse
		CleanSearchHTML = originalCleanSearchHTML
	}()

	// Replace originals with mocks
	getSearchResponse = mockGetSearchResponse
	CleanSearchHTML = mockCleanSearchHTML

	tests := []struct {
		query    string
		expected string
		err      bool
	}{
		{"valid query", "Search result", false},
		{"empty query", "", true},
		{"invalid query", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			result, err := WikiSearch(tt.query)
			if (err != nil) != tt.err {
				t.Errorf("expected err: %v, got: %v", tt.err, err)
			}
			if result != tt.expected {
				t.Errorf("expected: %s, got: %s", tt.expected, result)
			}
		})
	}
}
