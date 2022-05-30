package main

import (
	"fmt"
	"io/ioutil"
)

func ifelse() {
	const filename = "/Users/huxiaoting01/Coding/go/learngo/go-learning/3_Conditions/demo.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err) //open demo.txt: no such file or directory
	//} else {
	//	fmt.Printf("%s \n", contents)
	//}

	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 注意：在 if 之外的block，不能打印 contents 和 err
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}

func switchdemo(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func grade(score int) string {
	desc := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		desc = "F"
	case score < 80:
		desc = "C"
	case score < 90:
		desc = "B"
	case score <= 100:
		desc = "A"
	}
	return desc
}

func main() {
	ifelse()
	add := switchdemo(1, 2, "+")
	sub := switchdemo(1, 2, "-")
	//a := switchdemo(1, 2, "<")

	fmt.Println(add)
	fmt.Println(sub)
	//fmt.Println(a)

	fmt.Println(grade(98))
}
