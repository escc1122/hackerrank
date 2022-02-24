package main

import "fmt"

func Example_matchingStrings() {
	strings := []string{"-4", "-4", "-9", "0"}
	queries := []string{"-4", "3", "-9", "0"}
	fmt.Println(matchingStrings(strings, queries))
	// Output:
	// [2 0 1 1]

}
