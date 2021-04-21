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