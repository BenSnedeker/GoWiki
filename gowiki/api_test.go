package gowiki

import "testing"

// General testing while the function is being built
func TestWikiSearch(t *testing.T) {
	// Establish test flags
	flags := Flags{
		Term:    "Gourd",
		Type:    ReturnType(Text),
		Content: ReturnContent(All),
		Style:   ReturnStyle(Fancy),
	}

	// Attempt WikiSearch
	response, err := WikiSearch(&flags)
	if err != nil {
		t.Fatalf("WikiSearch failed to search Wikipedia: %v", err)
	}

	if len(response) < 1 {
		// Perhaps fail if no results come back?
	}
}

// Test WikiSearch for an empty Term
// Test WikiSearch for an HTML search
// Test WikiSearch for a Text search
// Test WikiSearch for a non-sense search
// Test WikiSearch for a clean response

// Test WikiRead for an empty Term
// Test WikiRead for an HTML response
// Test WikiRead for a text response
// Test WikiRead for a non-sense article
// Test WikiRead for a suggestion page
// Test WikiRead for All sections
// Test WikiRead for Limited sections
// Test WikiRead for Summary section
// Test WikiRead for References section
// Test WikiRead for a fancy response
// Test WikiRead for a clean response
