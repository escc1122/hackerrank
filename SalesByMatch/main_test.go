package main

import "fmt"

func Example_sockMerchant() {
	test := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(len(test)), test))
	// Output:
	// 3
}
