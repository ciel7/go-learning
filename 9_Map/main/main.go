package main

import "fmt"

// main
// map[K]V
// map[K1]map[K2]V
// map 不保证遍历顺序，如需排序，可以将其一个个加入到slice，slice可以排序
// map 使用哈希表，必须可以比较相等
// 除了 slice,map,function的内建类型都可以作为 key
// struct 类型不包含上述字段时，也可以作为key
func mapDefine() map[string]string {
	m := map[string]string{
		"name": "tutu",
		"age":  "18",
	}

	return m
}
func mapDefine1() map[string]int {
	m := make(map[string]int) // m == empty map
	m = map[string]int{
		"min": 5,
		"max": 10,
	}

	return m
}

func mapDefine2() map[string]int {
	var m map[string]int // m == nil
	m = map[string]int{
		"min": 5,
		"max": 10,
	}

	return m
}
func main() {
	m := mapDefine()
	fmt.Println(m)

	m1 := mapDefine1()
	fmt.Println(m1)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println("k = ", k)
		fmt.Println("v = ", v)
	}

	fmt.Println("Getting Values")
	min, ok := m1["min"]
	fmt.Println("min = ", min)
	fmt.Println("min ok = ", ok) // true

	if minx, ok := m1["minx"]; ok {
		fmt.Println("minx = ", minx)  // 获取 map 中不存在的key的元素时，会返回该类型的初始值
		fmt.Println("minx ok = ", ok) // 取 map[string]int 中  不存在的key时，false
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting Values")

	delete(m1, "max")
	fmt.Println("m1 = ", m1)

}
