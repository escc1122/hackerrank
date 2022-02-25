package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	returnValue := "pangram"
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	// var countMap map[byte]int
	countMap := make(map[rune]int, 26)
	for i := 97; i <= 122; i++ {
		countMap[rune(i)] = 0
	}

	for i := 0; i < len(s); i++ {
		countMap[rune(s[i])] = 1
	}
	for _, v := range countMap {
		if v == 0 {
			returnValue = "not pangram"
			break
		}
	}
	return returnValue
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
