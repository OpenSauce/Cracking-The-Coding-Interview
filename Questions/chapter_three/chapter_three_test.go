package chapter_three

import (
	"testing"
)

func TestQuestionOne(t *testing.T) {
	//Not a function question
}

func TestQuestionTwo(t *testing.T) {
	//Not a function question
}

func TestQuestionThree(t *testing.T) {
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
		got := QuestionThree(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}
	}
}
