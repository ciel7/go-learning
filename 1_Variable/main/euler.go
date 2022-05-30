package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// euler 欧拉公式
func cmplxText() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

// euler 欧拉公式
// 复数的实部和虚部都是用 float 来表示的
// complex64：实部和虚部分别是 float32
// complex128：实部和虚部分别是 float64
func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // (0+1.2246467991473515e-16i)
	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) //(0.000+0.000i)
}

func main() {
	cmplxText()
	euler()
}
