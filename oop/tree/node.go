package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("can't set Value to nil")
		return
	}
	node.Value = Value
}




