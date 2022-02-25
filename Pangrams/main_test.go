package main

import "fmt"

func Example_pangrams() {
	test := "We promptly judged antique ivory buckles for the prize"
	fmt.Println(pangrams(test))
	// Output:
	// not pangram
}
