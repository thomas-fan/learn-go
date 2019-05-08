package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverserFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverserFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverserFunc(f)
	f(node)
	node.Right.TraverserFunc(f)
}
