package main

import "fmt"

// 函数外部的变量不是全局变量，他们属于 <<< 包内变量 >>>
var aa = 3
var bb = "bbb"

// 这种写法是不被允许的
// 函数外面一定要以关键字开头
//bb := "abc"

var (
	cc = 77
	dd = "qaz"
)

/**
variableZeroValue 变量默认初值
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("a = %d, s = %q \n", a, s)
}

/**
variableTypeDeduction 变量初始化赋值
*/
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("a = %d, b = %d, s = %q \n", a, b, s)
}

/**
variableTypeDeduction 类型推断
*/
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

/**
variableTypeDeductionShorter 变量初始化赋值
*/
func variableTypeDeductionShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableTypeDeductionShorter()

	fmt.Println(aa, bb, cc, dd)
}
