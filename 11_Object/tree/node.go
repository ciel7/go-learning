package tree

import "fmt"

// Node 结构体是创建在堆上还是栈上？
// C++：局部变量是分配在栈上的，函数一旦结束，这个局部变量就会被销毁；如果想传递出去，就需要在堆上创建，堆上分配就需要手动释放
// Java：几乎所有东西都是分配在堆上，有垃圾回收机制
// 对于 Go 语言来说，开发者不需要知道是分配在堆上还是栈上:
// 因为 Go 有垃圾回收器; 这是由 Go 语言编译器 和 运行环境决定的。
// 比如说：Node 这个局部变量没有取地址，很可能就会被认为这个变量不需要给外部使用，就创建在栈中；如果发现你取地址了并返回给别人用了，那么就创建在堆中。
// 当 Node 分配在堆中，就会参与垃圾回收。
type Node struct {
	Value       int
	Left, Right *Node
}

// Print 为结构体定义方法
// (node Node) 称为接收者 - 值传递
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue 如果直接这样写，SetValue 并不能改变 node 的值，因为是值传递！
//func (node Node) SetValue(Value int) {
//	node.Value = Value
//}
// SetValue 将接收者改成指针类型
// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored. ")
		return
	}
	node.Value = Value
}

// CreatNode go 没有构造函数，如果想把创建放在我们这里控制，可以增加工厂函数
func CreatNode(Value int) *Node {
	// 注意 Go 语言返回局部变量的地址，在这里不会报错
	return &Node{Value: Value}
}
