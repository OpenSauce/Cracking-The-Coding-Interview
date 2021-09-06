package chapter_one

import (
	"reflect"
	"testing"
)

func TestQuestionOne(t *testing.T) {
	t.Run("AAAA", func(t *testing.T) {
		got := QuestionOne("AAAAAA")
		want := false

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("ABCDEF", func(t *testing.T) {
		got := QuestionOne("ABCDEFGH")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestQuestionTwo(t *testing.T) {
	t.Run("Permutation", func(t *testing.T) {
		got := QuestionTwo("CAT", "TAC")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("Not a permutation", func(t *testing.T) {
		got := QuestionTwo("Bob", "Flubber")
		want := false

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestQuestionThree(t *testing.T) {
	t.Run("URLify string one", func(t *testing.T) {
		got := QuestionThree("Mr John Smith    ", 13)
		want := "Mr%20John%20Smith"

		if got != want {
			t.Errorf("got: %s; want: %s", got, want)
		}
	})

	t.Run("URLify string two", func(t *testing.T) {
		got := QuestionThree("One Two Three Four      ", 18)
		want := "One%20Two%20Three%20Four"

		if got != want {
			t.Errorf("got: %s; want: %s", got, want)
		}
	})
}

func TestQuestionFour(t *testing.T) {
	t.Run("Palindrome #1", func(t *testing.T) {
		got := QuestionFour("tactcoa")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("Panlindrome #2", func(t *testing.T) {
		got := QuestionFour("bobohehe")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("Panlindrome #3", func(t *testing.T) {
		got := QuestionFour("hellothere")
		want := false

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestQuestionFive(t *testing.T) {
	t.Run("One Away #1", func(t *testing.T) {
		got := QuestionFive("Hello", "Hello1")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("One Away #2", func(t *testing.T) {
		got := QuestionFive("Bake", "Pake")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("One Away #3", func(t *testing.T) {
		got := QuestionFive("One", "One")
		want := true

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestQuestionSix(t *testing.T) {
	var testCase = []struct {
		in   string
		want string
	}{
		{"aaaabbbbcccddddeff", "a4b4c3d4e1f2"},
		{"aaaabbbbcccddddef", "a4b4c3d4e1f1"},
		{"aaaabbbbcccddddefgg", "a4b4c3d4e1f1g2"},
		{"abcdefg", "a1b1c1d1e1f1g1"},
	}

	for _, tt := range testCase {
		got := QuestionSix(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}

func TestQuestionSeven(t *testing.T) {
	var testCase = []struct {
		in   [][]int
		want [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			[][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			},
		},
		{
			[][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			},
			[][]int{
				{3, 2, 1},
				{3, 2, 1},
				{3, 2, 1},
			},
		},
		{
			[][]int{
				{1, 1, 1, 1},
				{2, 2, 2, 2},
				{3, 3, 3, 3},
				{4, 4, 4, 4},
			},
			[][]int{
				{4, 3, 2, 1},
				{4, 3, 2, 1},
				{4, 3, 2, 1},
				{4, 3, 2, 1},
			},
		},
	}

	for _, tt := range testCase {
		got := QuestionSeven(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}

func TestQuestionEight(t *testing.T) {
	var testCase = []struct {
		in   [][]int
		want [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
		},
		{
			[][]int{
				{1, 1, 1},
				{0, 2, 2},
				{3, 3, 3},
				{4, 4, 4},
			},
			[][]int{
				{0, 1, 1},
				{0, 0, 0},
				{0, 3, 3},
				{0, 4, 4},
			},
		},
		{
			[][]int{
				{1, 1, 1, 1},
				{2, 2, 2, 2},
				{3, 3, 0, 3},
				{4, 4, 4, 4},
				{5, 5, 5, 5},
			},
			[][]int{
				{1, 1, 0, 1},
				{2, 2, 0, 2},
				{0, 0, 0, 0},
				{4, 4, 0, 4},
				{5, 5, 0, 5},
			},
		},
	}

	for _, tt := range testCase {
		got := QuestionEight(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}

func TestQuestionNine(t *testing.T) {
	var testCase = []struct {
		in   string
		in2  string
		want bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbottle", "erbottlemate", false},
		{"hellomynameis", "nameishellomy", true},
		{"tabletennis", "pingpong", false},
	}

	for _, tt := range testCase {
		got := QuestionNine(tt.in, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}
