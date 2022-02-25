package main

import "fmt"

func Example_twoArrays() {
	A := []int32{2, 1, 3}
	B := []int32{7, 8, 9}
	var k int32 = 10
	fmt.Println(twoArrays(k, A, B))
	// Output:
	// YES
}
