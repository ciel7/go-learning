package main

import "fmt"

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
}
