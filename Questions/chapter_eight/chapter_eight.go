package chapter_eight

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
