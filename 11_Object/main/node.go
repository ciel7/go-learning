package main

import "fmt"

// 结构体是创建在堆上还是栈上？
// C++：局部变量是分配在栈上的，函数一旦结束，这个局部变量就会被销毁；如果想传递出去，就需要在堆上创建，堆上分配就需要手动释放
// Java：几乎所有东西都是分配在堆上，有垃圾回收机制
// 对于 Go 语言来说，开发者不需要知道是分配在堆上还是栈上:
// 因为 Go 有垃圾回收器; 这是由 Go 语言编译器 和 运行环境决定的。
// 比如说：treeNode 这个局部变量没有取地址，很可能就会被认为这个变量不需要给外部使用，就创建在栈中；如果发现你取地址了并返回给别人用了，那么就创建在堆中。
// 当 treeNode 分配在堆中，就会参与垃圾回收。
type treeNode struct {
	value       int
	left, right *treeNode
}

// 为结构体定义方法
// (node treeNode) 称为接收者 - 值传递
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 如果直接这样写，setValue 并不能改变 node 的值，因为是值传递！
//func (node treeNode) setValue(value int) {
//	node.value = value
//}
// setValue 将接收者改成指针类型
// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored. ")
		return
	}
	node.value = value
}

// go 没有构造函数，如果想把创建放在我们这里控制，可以增加工厂函数
func creatNode(value int) *treeNode {
	// 注意 Go 语言返回局部变量的地址，在这里不会报错
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	//var root treeNode
	//fmt.Println(root) // {0 <nil> <nil>}

	root := treeNode{
		value: 3,
	}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // new(treeNode) 返回的是一个地址
	root.left.right = creatNode(2)

	//root.print()
	root.right.left.setValue(4)
	//root.right.left.print()

	root.traverse()

	//var pRoot *treeNode
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	// 在 slice 中 treeNode 可以省略
	//nodes := []treeNode{
	//	//treeNode{value: 3},
	//	//treeNode{},
	//	//treeNode{6, nil, &root},
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
}

// 值接收者 vs 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也考虑指针接收者
// 一致性：如有指针接收者，最好都是指针接收者

// 值接收者是 go 语言特有
// 值/指针接收者均可接收值/指针
