package main

import "fmt"

func swapVal(a, b int) {
	a, b = b, a
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// go 语言只有一种值传递方式，不支持引用传递
func main() {
	a := 3
	b := 4
	swapVal(a, b)
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}
