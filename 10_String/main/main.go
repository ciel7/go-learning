package main

import (
	"fmt"
	"unicode/utf8"
)

// main rune 相当于 go 的 char
// 可以使用 range 遍历 pos, rune 对（pos在中文字符的时候是不连续的）
// 可以使用 utf8.RuneCountInString 获得字符的数量
// 使用 len 可以获得字节的长度
// 使用 []byte 可以获得字节
//
func main() {
	s := "Hello我是图图!" // UTF-8
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 48 65 6C 6C 6F E6 88 91 E6 98 AF E5 9B BE E5 9B BE 21
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf(" (%d %X) ", i, ch) //  (0 48)  (1 65)  (2 6C)  (3 6C)  (4 6F)  (5 6211)  (8 662F)  (11 56FE)  (14 56FE)  (17 21)
	}
	fmt.Println()

	fmt.Println("Runr Count:", utf8.RuneCountInString(s)) //Runr Count: 10

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("ch = %c \n", ch) //rune = 72,size = 1
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("i = %d, ch = %c \n", i, ch)
	}

	fmt.Println()
}
