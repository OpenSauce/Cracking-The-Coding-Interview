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
