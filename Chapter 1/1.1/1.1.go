package main

import (
	"sort"
	"strings"
)

func QuestionOne(word string) bool {
	var m = make(map[rune]int)

	for _, char := range word {
		if m[char] == 1 {
			return false
		}
		m[char] = 1
	}
	return true
}

func QuestionTwo(word1 string, word2 string) bool {
	s := strings.Split(word1, "")
	sort.Strings(s)
	s2 := strings.Split(word2, "")
	sort.Strings(s2)
	return strings.Join(s, "") == strings.Join(s2, "")
}

func QuestionThree(word string, length int) string {
	return ""
}