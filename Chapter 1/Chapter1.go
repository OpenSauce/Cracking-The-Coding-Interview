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
	var res string
	for index := 0; index < length; index++ {
		if word[index] == ' ' {
			res += "%20"
		} else {
			res += string(word[index])
		}
	}
	return res
}

func QuestionFour(word string) bool {
	var m = make(map[rune]int)

	for _, char := range word {
		if m[char] == 1 {
			delete(m, char)
		} else {
			m[char] = 1
		}
	}

	return len(m) <= 1
}

func QuestionFive(word1 string, word2 string) bool {
	var m = make(map[rune]int)

	for _, char := range word1 {
		if m[char] == 1 {
			delete(m, char)
		} else {
			m[char] = 1
		}
	}

	for _, char := range word2 {
		if m[char] == 1 {
			delete(m, char)
		} else {
			m[char] = 1
		}
	}
	return true
}