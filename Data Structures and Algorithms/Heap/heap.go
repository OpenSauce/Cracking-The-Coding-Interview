package main

import "fmt"

var tree_array_size = 11
var heap_size = 10

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func GetRightChild(A []int, index int) int {
	if (((2 * index) + 1) < tree_array_size) && (index >= 1) {
		return (2 * index) + 1
	}
	return -1
}

func GetLeftChild(a []int, index int) int {
	if ((2 * index) < tree_array_size) && (index >= 1) {
		return 2 * index
	}
	return -1
}

func GetParent(A []int, index int) int {
	if (index > 1) && (index < tree_array_size) {
		return index / 2
	}
	return -1
}

func max_heapify(A []int, index int) {
	left_child_index := GetLeftChild(A, index)
	right_child_index := GetRightChild(A, index)

	largest := index

	if (left_child_index <= heap_size) && (left_child_index > 0) {
		if A[left_child_index] > A[largest] {
			largest = left_child_index
		}
	}

	if right_child_index <= heap_size && (right_child_index > 0) {
		if A[right_child_index] > A[largest] {
			largest = right_child_index
		}
	}

	// largest is not the node, node is not a heap
	if largest != index {
		swap(&A[index], &A[largest])
		max_heapify(A, largest)
	}
}

func build_max_heap(A []int) {
	for i := heap_size / 2; i >= 1; i-- {
		max_heapify(A, i)
	}
}

func main() {
	var A = []int{0, 15, 20, 7, 9, 5, 8, 6, 10, 2, 1}
	build_max_heap(A)
	for i := 1; i <= heap_size; i++ {
		fmt.Printf("%d\n", A[i])
	}
}