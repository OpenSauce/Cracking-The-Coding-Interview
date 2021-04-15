package main

import (
	"fmt"
)


type Node struct {
	data   int
	right  *Node
	left   *Node
	parent *Node
}

type Tree struct {
	root *Node
}

func NewNode(data int) *Node {
	n := &Node{
		data:   data,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	return n
}

func NewTree() *Tree {
	t := &Tree{
		root: nil,
	}

	return t
}

func Minimum(t *Tree, x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func Insert(t *Tree, n *Node) {
	var y *Node
	temp := t.root
	for temp != nil {
		y = temp
		if n.data < temp.data {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	n.parent = y

	if y == nil {
		t.root = n
	} else if n.data < y.data {
		y.left = n
	} else {
		y.right = n
	}
}

func Transplant(t *Tree, u *Node, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func Delete(t *Tree, z *Node) {
	if z.left == nil {
		Transplant(t, z, z.right)
	} else if z.right == nil {
		Transplant(t, z, z.left)
	} else {
		y := Minimum(t, z.right) 
		if y.parent != z {
			Transplant(t, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		Transplant(t, z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func Inorder(t *Tree, n *Node) {
	if n != nil {
		Inorder(t, n.left)
		fmt.Printf("%d\n", n.data)
		Inorder(t, n.right)
	}
}

func main() {
	t := NewTree()

	a := NewNode(10)
	b := NewNode(20)
	c := NewNode(30)
	d := NewNode(100)
	e := NewNode(90)
	f := NewNode(40)
	g := NewNode(50)
	h := NewNode(60)
	i := NewNode(70)
	j := NewNode(80)
	k := NewNode(150)
	l := NewNode(110)
	m := NewNode(120)

	Insert(t, a)
	Insert(t, b)
	Insert(t, c)
	Insert(t, d)
	Insert(t, e)
	Insert(t, f)
	Insert(t, g)
	Insert(t, h)
	Insert(t, i)
	Insert(t, j)
	Insert(t, k)
	Insert(t, l)
	Insert(t, m)

	Delete(t, a)
	Delete(t, m)

	Inorder(t, t.root)
}