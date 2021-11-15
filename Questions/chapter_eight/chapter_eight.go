package chapter_eight

import "fmt"

func QuestionOne(n int) int {
	memo := make([]int, n+1)
	for i, _ := range memo {
		memo[i] = -1
	}
	return countWays(n, memo)
}

func countWays(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if memo[n] > -1 {
		return memo[n]
	} else {
		memo[n] = countWays(n-1, memo) + countWays(n-2, memo) + countWays(n-3, memo)
		return memo[n]
	}
}

type point struct {
	x int
	y int
}

func QuestionTwo(grid [][]int) []int {
	var path []int
	failedPoints := make(map[point]struct{})
	if getPath(grid, len(grid[0]), len(grid), path, failedPoints) {
		return path
	} else {
		return nil
	}
}

func getPath(grid [][]int, row int, column int, path []int, failedPoints map[point]struct{}) bool {
	if grid[row][column] == 0 {
		return false
	}

	_, ok := failedPoints[point{row, column}]
	if ok {
		return false
	}

	atOrigin := row == 0 && column == 0

	if atOrigin || getPath(grid, row, column-1, path, failedPoints) || getPath(grid, row-1, column, path, failedPoints) {
		path = append(path, grid[row][column])
		return true
	}

	failedPoints[point{row, column}] = struct{}{}
	return false
}

func QuestionThree(values []int) int {
	return BinarySearch(values, 0, len(values)-1)
}

func BinarySearch(values []int, left int, right int) int {
	if left <= right {
		mid := (left + right) / 2
		if values[mid] == mid {
			return mid
		}

		if values[mid] > mid {
			return BinarySearch(values, left, mid-1)
		}

		if values[mid] < mid {
			return BinarySearch(values, mid+1, right)
		}
	}
	return -1
}

func QuestionFour(set []int) [][]int {
	var result [][]int
	for i := 0; i < (1 << len(set)); i++ {
		var newVal []int
		for j := 0; j < len(set); j++ {
			if (i & (1 << j)) > 0 {
				newVal = append(newVal, set[j])
			}
		}
		fmt.Printf("%+v\n", newVal)
		result = append(result, newVal)
	}
	return result
}

func QuestionFive(num1 int, num2 int) int {
	c := 0
	t := num1
	for (t ^ 1) == t+1 {
		t >>= 1
		c++
	}
	res := multiply(t, num2)
	return res << c
}

func multiply(smaller int, bigger int) int {
	if smaller == 0 {
		return 0
	}
	if smaller == 1 {
		return bigger
	}

	s := smaller >> 1
	halfProd := multiply(s, bigger)

	if smaller%2 == 0 {
		return halfProd << 1
	} else {
		return (halfProd << 1) + bigger
	}
}

type Stack struct {
	top      *StackNode
	capacity int
}

type StackNode struct {
	data interface{}
	next *StackNode
}

func (s *Stack) pop() (interface{}, error) {
	if s.top != nil {
		data := s.top.data
		s.top = s.top.next
		return data, nil
	}
	return 0, fmt.Errorf("empty Stack")
}

func (s *Stack) push(in interface{}) {
	t := StackNode{data: in, next: s.top}
	s.top = &t
}

func (s *Stack) peek() (interface{}, error) {
	if s.top != nil {
		return s.top.data, nil
	}
	return 0, fmt.Errorf("empty Stack")
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func MoveTo(n int, origin Stack, dest Stack, buffer Stack) {
	if n == 0 {
		return
	}
	MoveTo(n-1, origin, buffer, dest)
	val, _ := origin.pop()
	dest.push(val)
	MoveTo(n-1, buffer, dest, origin)
}

// Permuatations of a string
func GetPerms(characters string) []string {
	var result []string
	if len(characters) == 0 {
		result = append(result, "")
		return result
	}

	for i := 0; i < len(characters); i++ {
		before := characters[:i]
		after := characters[i+1:]
		perms := GetPerms(before + after)

		for _, v := range perms {
			fmt.Println(string(characters[i]) + v)
			result = append(result, string(characters[i])+v)
		}
	}
	return result
}
