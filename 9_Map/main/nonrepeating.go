package main

import "fmt"

// main
// 寻找最长不含有重复字符的子串
// abcabcbb -> abc
// bbbbb -> b
// pwwkew -> wke

// lengthOfNonRepeatingSubStr
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	// 遍历字符串
	for i, ch := range []byte(s) {
		fmt.Println("ch = ", ch)
		lastI, ok := lastOccurred[ch]
		fmt.Println("ch 出现的下标 = ", lastI)
		fmt.Println("ch ok = ", ok)

		if ok && lastI >= start {
			// 更新 start
			start = lastI + 1
			fmt.Println("ch 出现的下标 >= start, start = lastI + 1", start)
		}
		fmt.Println("当前最长子串长度 = ", maxLength)
		fmt.Println("当前 ch 距离 start 的长度 = ", i-start+1)
		// i 到 start 的元素个数
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			fmt.Println("maxLength 更新 = ", i-start+1)
		}

		lastOccurred[ch] = i
		fmt.Println("lastOccurred = ", lastOccurred)
	}

	return maxLength
}

// 目前该算法不支持中文，后面学到 rune 时会处理
// 算法概述：
// 对于每一个字母 x
// lastOccurred[x] 不存在，或者 < start: 无需操作
// lastOccurred[x] >= start: 更新 start
// 更新 lastOccurred[x]，更新 maxLength

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	//fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	//fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	//fmt.Println(lengthOfNonRepeatingSubStr(""))
	//fmt.Println(lengthOfNonRepeatingSubStr("b"))
	//fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
}
