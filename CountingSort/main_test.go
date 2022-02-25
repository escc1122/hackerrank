package main

import "fmt"

func Example_diagonalDifference() {
	test := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	fmt.Println(diagonalDifference(test))
	// Output:
	// 15

}
