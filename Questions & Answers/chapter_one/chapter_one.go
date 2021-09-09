package chapter_one

import (
	"fmt"
	"sort"
	"strings"
)

// func QuestionOne(word string) bool {
// 	var m = make(map[rune]int)

// 	for _, char := range word {
// 		if m[char] == 1 {
// 			return false
// 		}
// 		m[char] = 1
// 	}
// 	return true
// }

func QuestionOne(word string) bool {
	checker := 0
	for _, char := range word {
		val := char - 'a'
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
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
	var m = make(map[rune]struct{})

	for _, char := range word1 {
		if m[char] == struct{}{} {
			delete(m, char)
		} else {
			m[char] = struct{}{}
		}
	}

	for _, char := range word2 {
		if m[char] == struct{}{} {
			delete(m, char)
		} else {
			m[char] = struct{}{}
		}
	}

	return len(m) <= 1 || (len(m) == 2 && len(word2) == len(word1))
}

func QuestionSix(input string) string {
	var res string
	for i := 0; i < len(input); i++ {
		char := input[i]
		for j := i; j < len(input); j++ {
			if char == input[j] && j+1 != len(input) {
				continue
			}

			val := j - i
			if j+1 == len(input) {
				if char != input[j] {
					res += fmt.Sprintf("%c%d", char, val)
					res += fmt.Sprintf("%c%d", input[j], val)
					return res
				}
				res += fmt.Sprintf("%c%d", char, val+1)
				return res
			}

			res += fmt.Sprintf("%c%d", char, val)
			i = j - 1

			break
		}
	}
	return res
}

func QuestionSeven(in [][]int) [][]int {
	res := make([][]int, len(in))
	for i := range res {
		res[i] = make([]int, len(in))
	}

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			x := j
			y := len(in) - i - 1
			res[x][y] = in[i][j]
		}
	}
	return res
}

func QuestionEight(in [][]int) [][]int {
	res := in

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == 0 {
				for k := 0; k < len(in); k++ {
					res[k][j] = 0
				}
				for k := 0; k < len(in[i]); k++ {
					res[i][k] = 0
				}
				return res
			}
		}
	}
	return res
}

func QuestionNine(in, in2 string) bool {
	fullString := in + in
	return strings.Contains(fullString, in2)
}
