package chapter_ten

import "sort"

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
