package main

import "fmt"

func Example_birthday() {
	test := []int32{1, 2, 1, 3, 2}
	fmt.Println(birthday(test, 3, 2))
	// Output:
	// 2
}
