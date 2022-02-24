package main

import "fmt"

func Example_timeConversion() {
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("12:40:22AM"))
	fmt.Println(timeConversion("12:45:54PM"))
	// Output:
	// 19:05:45
	// 00:40:22
	// 12:45:54

}
