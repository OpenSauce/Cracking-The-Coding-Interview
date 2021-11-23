package chapter_ten

import (
	"fmt"
	"testing"
)

func TestSortedMerge(t *testing.T) {
	fmt.Println(SortedMerge([]int{2, 5, 6}, []int{4, 7, 8}))
}

func TestSearchRotatedArray(t *testing.T) {
	fmt.Println(SearchRotatedArray([]int{10, 12, 15, 2, 4, 6, 8}, 8))
}
