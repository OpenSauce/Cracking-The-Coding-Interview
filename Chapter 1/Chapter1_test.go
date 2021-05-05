package main

import "testing"

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
		got := QuestionFive("Hello", "Hellol")
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
		got := QuestionFive("One", "Two")
		want := false

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}
