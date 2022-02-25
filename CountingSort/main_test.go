package main

import "fmt"

func Example_countingSort() {
	test := []int32{11, 2, 4, 6, 7, 2, 22, 99}
	fmt.Println(countingSort(test))
	// Output:
	// [0 0 2 0 1 0 1 1 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

}
