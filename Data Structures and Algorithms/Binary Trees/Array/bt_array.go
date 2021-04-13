package main

import "fmt"

var complete_node = 15
var tree = []byte{'0', 'D', 'A', 'F', 'E', 'B', 'R', 'T', 'G', 'Q', '0', '0', 'V', '0', 'J', 'L'}

func GetRightChild(index int) int {
	if tree[index] != '0' && ((2*index)+1) <= complete_node {
		return (2 * index) + 1
	}
	return -1
}

func GetLeftChild(index int) int {
	if tree[index] != '0' && (2*index) <= complete_node {
		return 2 * index
	}
	return -1
}

func GetParent(index int) int {
	if tree[index] != '0' && index/2 != '0' {
		return index / 2
	}
	return -1
}

func Preorder(index int) {

	if index > 0 && tree[index] != '0' {
		fmt.Printf(" %c ", tree[index]) // visiting root
		Preorder(GetLeftChild(index))   //visiting left subtree
		Preorder(GetLeftChild(index))   //visiting right subtree
	}
}

func Postorder(index int) {

	if index > 0 && tree[index] != '0' {
		Postorder(GetLeftChild(index)) //visiting left subtree
		Postorder(GetLeftChild(index)) //visiting right subtree
		fmt.Printf(" %c ", tree[index])    //visiting root
	}
}

func Inorder(index int) {

	if index > 0 && tree[index] != '0' {
		Inorder(GetLeftChild(index)) //visiting left subtree
		fmt.Printf(" %c ", tree[index])  //visiting root
		Inorder(GetLeftChild(index)) // visiting right subtree
	}
}

func main() {
	fmt.Printf("Preorder:\n")
	Preorder(1)
	fmt.Printf("\nPostorder:\n")
	Postorder(1)
	fmt.Printf("\nInorder:\n")
	Inorder(1)
	fmt.Printf("\n")
}