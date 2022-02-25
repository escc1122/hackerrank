package main

import "fmt"

func Example_flippingBits() {

	fmt.Println(flippingBits(2147483647))
	fmt.Println(flippingBits(1))
	fmt.Println(flippingBits(0))

	// Output:
	// 2147483648
	// 4294967294
	// 4294967295

}
