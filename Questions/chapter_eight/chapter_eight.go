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
