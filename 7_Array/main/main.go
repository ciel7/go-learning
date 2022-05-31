package main

import "fmt"

// printArray 数组要指定长度，有些麻烦，go语言一般使用切片
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 数组是值类型传递，这里并不会影响main函数中数组的值
	arr[0] = 100
}

// main 数组是值类型
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

	for i := 0; i < len(grid); i++ {

		// range
		// 意义明确，美观
		//
		for j, cont := range grid[i] {
			fmt.Println(j)
			fmt.Println(cont)
		}
	}

	// 注意：这里的数组只能是[5]int的，如果是其他数字个数的int数组，就会报错
	// go语言中，[10]int 和 [20]int 是不同的类型
	// 调用 func f(arr [10]int) 会拷贝数组
	arr5 := [5]int{2, 4, 6, 8, 0}
	printArray(arr5)
}
