package main

import "fmt"

func Example_lonelyinteger() {
	test := []int32{1, 1, 3, 2, 2}
	fmt.Println(lonelyinteger(test))
	// Output:
	// 3

}
