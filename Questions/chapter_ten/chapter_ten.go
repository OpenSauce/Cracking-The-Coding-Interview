package chapter_ten

import (
	"sort"
	"strings"
)

func SortedMerge(array []int, array2 []int) []int {
	result := make([]int, len(array)+len(array2))
	indexA := len(array) - len(array2) + 2
	indexB := len(array2) - 1

	for i := len(result) - 1; i >= 0; i-- {
		if indexB == -1 || array[indexA] > array2[indexB] {
			result[i] = array[indexA]
			indexA--
			continue
		}
		result[i] = array2[indexB]
		indexB--
	}

	return result
}

func GroupAnagrams(values []string) []string {
	mapList := make(map[string][]string)
	for _, s := range values {
		key := []rune(s)
		sort.Slice(key, func(i int, j int) bool { return key[i] < key[j] })
		mapList[string(key)] = append(mapList[string(key)], s)
	}

	var res []string
	for _, v := range mapList {
		res = append(res, v...)
	}
	return res
}

func SearchRotatedArray(values []int, x int) int {
	return search(values, x, 0, len(values)-1)
}

func search(values []int, x int, left int, right int) int {
	if right < left {
		return -1
	}

	mid := (left + right) / 2
	if x == values[mid] {
		return mid
	}
	if values[left] < values[mid] {
		if values[left] <= x && x < values[mid] {
			return search(values, x, left, mid-1)
		} else {
			return search(values, x, mid+1, right)
		}
	} else if values[mid] < values[right] {
		if values[mid] <= x && x <= values[right] {
			return search(values, x, mid+1, right)
		} else {
			return search(values, x, left, mid-1)
		}
	} else {
		location := -1
		if values[left] == values[mid] {
			location = search(values, x, mid+1, right)
		}

		if location == -1 && values[mid] == values[right] {
			location = search(values, x, left, mid-1)
		}
		return location
	}
}

/* Sorted search, no size

Binary search: but we don't know the size so can't start at half. So what is initial pivot? Maybe use the element as the size?
So element 100, we search at 50?


Use an exponential backoff. Use the initial value to see if the array is filled with max incremental values, then exponentially backoff until we find the value.

Find the length using an exponential increment, then use a standard binary search. The two O(log n) still counts as O(log n)
*/

func SparseSearch(vals []string, x string) int {
	return strSearch(vals, x, 0, len(vals)-1)
}

func strSearch(vals []string, x string, first, last int) int {
	if first > last {
		return -1
	}

	mid := (first + last) / 2

	if vals[mid] == "" {
		left := mid - 1
		right := mid + 1
		for {
			if left < right && right > last {
				return -1
			} else if right <= last && vals[right] != "" {
				mid = right
				break
			} else if left >= first && vals[left] != "" {
				mid = left
				break
			}
			right++
			left--
		}
	}

	switch strings.Compare(vals[mid], x) {
	case 1:
		return strSearch(vals, x, first, mid-1)
	case -1:
		return strSearch(vals, x, mid+1, last)
	default:
		return mid
	}
}

/* Sort Big File
Initially, some sort of sort where we don't need to iterate over the whole file, we can sort as we pass through. Space complexity is the most important thing here. Which would lead us to either a bubble sort or a selection sort.

Merge sort: Will need to copy the elements somewhere else, with a large file this is useful as we can partition the file into smaller sizes.

Divide the file into chunks, the size of the file being the amount of memory avaliable, we can sort each chunk seperately then merge theem in one by one.

External sort.

*/

/* Missing Int

Create a "bit vecor", an array of bits that are equal to the length of memory we have. We can then iterate through every value and set the bit of that value.

We can then iterate through that bit vector to find a value that has not been set to 1 and is 0.

*/
