package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 666
}

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

// slice 是array的一个view
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr[2:6] = ", arr[2:6]) // [2,3,4,5]
	fmt.Println("arr[2:] = ", arr[2:])   // [2,3,4,5,6,7,8,9]
	fmt.Println("arr[:6] = ", arr[:6])   // [0,1,2,3,4,5]
	fmt.Println("arr[:] = ", arr[:])     //  [0 1 2 3 4 5 6 7 8 9]

	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println("After UpdateSlice(s1)")
	fmt.Println("s1 = ", s1)
	fmt.Println("arr = ", arr)
	s2 := arr[:]
	fmt.Println("After UpdateSlice(s2)")
	updateSlice(s2)
	fmt.Println("s2 = ", s2)
	fmt.Println("arr = ", arr)

	s2 = s2[:5]
	fmt.Println("s2[:5] = ", s2)
	s2 = s2[2:]
	fmt.Println("s2[2:] = ", s2)
	s2 = s2[1:]
	fmt.Println("s2[1:] = ", s2)
	updateSlice(s2)
	fmt.Println("updateSlice s2 = ", s2)
	fmt.Println("updateSlice arr = ", arr)

	arr1 := [5]int{1, 2, 3, 4, 5}
	printArray(arr1[:])
	fmt.Println(arr1)

	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s3 := arr3[2:6] // [2,3,4,5]
	s4 := s3[3:5]   // 以为是 [5,0]，实际上是 [5 6]
	fmt.Println("s3[3] ", s3[3])
	// 会报错，因为此时 s3=[2,3,4,5]，那为啥 slice s1[3:5] 能取到值？就需要关注底层。
	//fmt.Println("s3[4] ", s3[4])
	fmt.Printf("s3 = %v, len(s3)=%d, cap(s3)=%d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4 = %v, len(s4)=%d, cap(s4)=%d \n", s4, len(s4), cap(s4))
}
