package gowiki

import "fmt"

// -------------
//    SEARCH
// -------------

func HandleSearch(query string) {
	fmt.Printf("Searching for \033[1;34m\"%s\"\033[0m\n", query)
}

// -------------
//     READ
// -------------

func HandleRead(flags Flags, query string) {
	fmt.Printf("Opening article \033[1;34m\"%s\"\033[0m\n", query)
}
