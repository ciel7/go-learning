package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	const (
		a, d = 3, 4
	)
	var c int
	// 这里和使用变量还不太一样，当我们 const a,b 没有指定类型的时候，在这里使用是没有问题的，变量就必须要做强制类型转换
	// 即：常量的数值可以作为各种类型使用
	c = int(math.Sqrt(a*a + d*d))
	fmt.Println(filename, c)

	//const (
	//	cpp        = 0
	//	java       = 1
	//	python     = 2
	//	golang     = 3
	//	javascript = 4
	//)

	// 上述等价于
	// iota 实现自增
	const (
		cpp = iota * 1024
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
}
