package chapter_ten

import (
	"fmt"
	"testing"
)

func TestSortedMerge(t *testing.T) {
	fmt.Println(SortedMerge([]int{2, 5, 6}, []int{4, 7, 8}))
}
