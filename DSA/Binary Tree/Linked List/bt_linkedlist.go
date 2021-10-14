package main

import "fmt"

type TreeNode struct {
	data   interface{}
	right  *TreeNode
	left   *TreeNode
	parent *TreeNode
}

func NewTreeNode(data interface{}) *TreeNode {
	n := &TreeNode{
		data:   data,
		right:  nil,
		left:   nil,
		parent: nil,
	}

	return n
}

type Tree struct {
	root *TreeNode
}

func NewTree(n *TreeNode) *Tree {
	t := &Tree{
		root: n,
	}

	return t
}

func Preorder(n *TreeNode) {
	if n != nil {
		fmt.Printf(" %c ", n.data)
		Preorder(n.left)
		Preorder(n.right)
	}
}

func Postorder(n *TreeNode) {
	if n != nil {
		Postorder(n.left)
		Postorder(n.right)
		fmt.Printf(" %c ", n.data)
	}
}

func Inorder(n *TreeNode) {
	if n != nil {
		Inorder(n.left)
		fmt.Printf(" %c ", n.data)
		Inorder(n.right)
	}
}

func main() {
	d := NewTreeNode('D')
	a := NewTreeNode('A')
	f := NewTreeNode('F')
	e := NewTreeNode('E')
	b := NewTreeNode('B')
	r := NewTreeNode('R')
	t1 := NewTreeNode('T')
	g := NewTreeNode('G')
	q := NewTreeNode('Q')
	v := NewTreeNode('V')
	j := NewTreeNode('J')
	l := NewTreeNode('L')

	t := NewTree(d)

	t.root.right = f
	t.root.left = a

	a.right = b
	a.left = e

	f.right = t1
	f.left = r

	e.right = q
	e.left = g

	r.left = v

	t1.right = l
	t1.left = j

	fmt.Printf("Preorder:\n")
	Preorder(t.root)
	fmt.Printf("\nPostorder:\n")
	Postorder(t.root)
	fmt.Printf("\nInorder:\n")
	Inorder(t.root)
	fmt.Printf("\n")
}
