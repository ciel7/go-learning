package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// div 带余除法 13 / 4 = 3 ... 1
func div(a, b int) (int, int) {
	return a / b, a % b
}

func div1(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数体较长时，不建议使用该方式
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div1(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling Function %s with args (%d, %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// sum 可变参数
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}
func main() {
	fmt.Println(div(13, 4))
	// cmd+alt+v 自动生成变量
	q, r := div1(13, 3)
	fmt.Println(q, r)

	q1, r1 := div2(13, 5)
	fmt.Println(q1, r1)

	//if i, err := eval(8, 9, "<"); err != nil {
	if i, err := eval(8, 9, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	a := 2
	b := 3
	fmt.Println(apply(pow, a, b))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, a, b))

	fmt.Println(
		sum(1, 2, 3),
	)
}
