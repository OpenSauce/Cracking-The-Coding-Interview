package main

type TreeNode struct {
	data   interface{}
	right  *TreeNode
	left   *TreeNode
	parent *TreeNode
}

func NewNode(data interface{}) *TreeNode {
	n := &TreeNode{
		data:   data,
		right:  nil,
		left:   nil,
		parent: nil,
	}

	return n
}

func main() {
}
s