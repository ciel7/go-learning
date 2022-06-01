package main

import (
	"go-learning/11_Object/tree"
)

// 包装实现自己的后序遍历
type myTreeNode struct {
	node *tree.Node
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
}

// 值接收者 vs 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也考虑指针接收者
// 一致性：如有指针接收者，最好都是指针接收者

// 值接收者是 go 语言特有
// 值/指针接收者均可接收值/指针
