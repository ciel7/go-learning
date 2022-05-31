package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3 = ", s3) // [5 6 10]
	// s4 and s5 no ;onger view arr.
	// s4 和 s4 的 view 不再是基于 arr，而是开辟了一个新的数组，是基于这个新的数组的view
	// 添加元素时，如果超越 cap，系统会重新分配更大的底层数组
	// 由于值传递的关系，必须接收 append 的返回值
	fmt.Println("s4 = ", s4)   // [5 6 10 11]
	fmt.Println("s5 = ", s5)   // [5 6 10 11 12]
	fmt.Println("arr = ", arr) // [0 1 2 3 4 5 6 10]
}
