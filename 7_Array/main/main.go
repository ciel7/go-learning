package main

import "fmt"

func main() {
	// 5个int组成的数组
	var arr1 [5]int
	// := 要给数组初值，所以要加 {}
	arr2 := [3]int{}
	// [...] 编译器会帮我们数有几个int
	arr3 := [...]int{}
	arr4 := [...]int{1, 3, 5, 7, 9}

	fmt.Println(arr1, arr2, arr3, arr4)

	var grid [4][5]int
	fmt.Println(grid)
}
