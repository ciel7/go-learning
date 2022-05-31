package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != err {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
		convertToBin(12345678978765),
	)
	filename := "/Users/huxiaoting01/Coding/go/learngo/go-learning/3_Conditions/demo.txt"
	printFile(filename)
}
