package gowiki

import (
	"testing"
)

// ------------------------
//   TESTS FOR WikiSearch
// ------------------------

// General testing while the function is being built
func TestWikiSearch(t *testing.T) {
	// Establish test flags
	flags := Flags{
		Term:     "Gourd",
		Type:     ReturnType(Text),
		Content:  ReturnContent(All),
		Style:    ReturnStyle(Fancy),
		Language: "en",
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

// Test WikiSearch for a malformed response (simulate HTTP error or invalid data)
// Test WikiSearch for an empty Term
// Test WikiSearch for a valid search term (common term)
// Test WikiSearch for a non-existent term (nonsense term)
// Test WikiSearch for a special character term (eg. $ % &)
// Test WikiSearch for a non-Latin characters (eg. Chinese Arabic Cyrillic)
// Test WikiSearch for an HTML response
// Test WikiSearch for a Text response
// Test WikiSearch for a clean response

// ------------------------
//    TESTS FOR WikiRead
// ------------------------

// Test WikiRead for a malformed response (simulate HTTP error or invalid data)
// Test WikiRead for an empty Term
// Test WikiRead for a valid search term (eg. Potato)
// Test WikiRead for a non-existent article (eg. jfdF34vS)
// Test WikiRead for a suggestion page (eg. Python)
// Test WikiRead for All sections
// Test WikiRead for Limited sections
// Test WikiRead for Summary section
// Test WikiRead for References section
// Test WikiRead for a clean response
// Test WikiRead for a fancy response
