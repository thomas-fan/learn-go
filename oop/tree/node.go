package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value, " ")
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("can't set value to nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse()  {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println()
	root.right.left.setValue(4)
	fmt.Println(root.right.left)
	root.traverse()
}
