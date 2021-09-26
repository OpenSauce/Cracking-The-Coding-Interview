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
	if root.Left != nil && root.Left.visited == false {
		search(root.Left)
	}
	if root.Right != nil && root.Right.visited == false {
		search(root.Right)
	}
}
