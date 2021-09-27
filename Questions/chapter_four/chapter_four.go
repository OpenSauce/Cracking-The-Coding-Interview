package chapter_four

import "fmt"

type Node struct {
	Data    int
	visited bool
	Right   *Node
	Left    *Node
}

type Tree struct {
	Root *Node
}

func QuestionOne(t1 *Tree, n1, n2 *Node) bool {
	search(t1.Root)

	return n1.visited && n2.visited
}

func search(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	root.visited = true
	if root.Left != nil && !root.Left.visited {
		search(root.Left)
	}
	if root.Right != nil && !root.Right.visited {
		search(root.Right)
	}
}

func QuestionTwo(input []int) *Tree {
	myTree := &Tree{}
	myTree.Root = InsertNode(input, 0, len(input)-1)
	return myTree
}

func InsertNode(input []int, start, end int) *Node {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	n := &Node{Data: input[mid]}
	n.Left = InsertNode(input, start, mid-1)
	n.Right = InsertNode(input, mid+1, end)
	return n
}
