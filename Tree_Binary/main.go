package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{Value: value, Left: nil, Right: nil}
}

func (node *TreeNode) Insert(val int, insertLeft bool) *TreeNode {
	if node == nil {
		return NewNode(val)
	}

	if insertLeft {
		node.Left = node.Left.Insert(val, !insertLeft)
	} else {
		node.Right = node.Right.Insert(val, !insertLeft)
	}
	return node
}

func (node *TreeNode) InOrderTraversal() {
	if node != nil {
		node.Left.InOrderTraversal()
		fmt.Print(node.Value, " ")
		node.Right.InOrderTraversal()
	}
}

func (node *TreeNode) GetRootNode() *TreeNode {
	return node
}

func main() {

	var root *TreeNode
	root = root.Insert(10, true)
	root.Insert(5, false)
	root.Insert(15, true)
	root.Insert(3, false)
	root.Insert(7, true)
	root.Insert(12, false)
	root.Insert(18, true)

	rootNode := root.GetRootNode()
	fmt.Println("Root Node :", rootNode.Value)

	root.InOrderTraversal()
}
