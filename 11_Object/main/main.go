package main

import (
	"fmt"
	"go-learning/11_Object/queue"
	"go-learning/11_Object/tree"
)

// 组合：包装实现自己的后序遍历
type myTreeNode struct {
	node *tree.Node
}

// 内嵌 Embedding（语法糖，省去部分代码量）
type myTreeNodeNew struct {
	*tree.Node
}

// 定义一个节点的后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func (myNode *myTreeNodeNew) postOrderNew() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNodeNew) Traverse() {
	fmt.Println("This method is shadowed.")
}

/**
	如何扩充系统类型或者别人的类型
	1. 组合：最常用
	2. 定义别名：最简单
    3. 内嵌 embedding：需要省下许多代码
*/
func main() {
	//var root Node
	//fmt.Println(root) // {0 <nil> <nil>}

	root := tree.Node{
		Value: 3,
	}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node) // new(Node) 返回的是一个地址
	root.Left.Right = tree.CreatNode(2)

	//root.Print()
	root.Right.Left.SetValue(4)
	//root.right.left.Print()

	root.Traverse()

	myroot := myTreeNode{&root}
	myroot.postOrder()

	root1 := myTreeNodeNew{&tree.Node{Value: 3}}
	root1.Left = &tree.Node{}
	root1.Right = &tree.Node{5, nil, nil}
	root1.Right.Left = new(tree.Node)
	root1.Left.Right = tree.CreatNode(2)
	root1.Right.Left.SetValue(4)
	fmt.Println("In-order traversal: ")
	// root1.Node.Traverse() 和 root1.Traverse() 虽然看似是继承，但实际上只是语法糖
	fmt.Println("root1.Node.Traverse()")
	root1.Node.Traverse() // 02345
	fmt.Println("root1.Traverse() ")
	root1.Traverse() // This method is shadowed.
	fmt.Print("My own post-order traversal: ")
	myroot1 := myTreeNodeNew{&root}
	myroot1.postOrderNew()
	fmt.Println()

	// 继承 和 go 语法糖的区别
	// 其他语言是可以把子类赋值给基类去使用的，但是在 go 中，没有继承这个概念，是不可以做到的
	//var baseRoot *tree.Node
	//baseRoot = &root
	// go 语言是通过接口来实现这样的能力的

	//var pRoot *Node
	//pRoot.SetValue(200)
	//pRoot = &root
	//pRoot.SetValue(300)
	//pRoot.Print()

	// 在 slice 中 Node 可以省略
	//nodes := []Node{
	//	//Node{value: 3},
	//	//Node{},
	//	//Node{6, nil, &root},
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	fmt.Println("测试 queue 部分")
	// 关注：11_Object/queue/queue.go
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}

// 值接收者 vs 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也考虑指针接收者
// 一致性：如有指针接收者，最好都是指针接收者

// 值接收者是 go 语言特有
// 值/指针接收者均可接收值/指针
