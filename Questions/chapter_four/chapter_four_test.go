package chapter_four

import "testing"

func TestQuestionOne(t *testing.T) {
	s1 := &Node{}
	s2 := &Node{}
	s3 := &Node{Left: s1}
	s4 := &Node{}
	s5 := &Node{Left: s4}
	s6 := &Node{Left: s3}
	s7 := &Node{Left: s5, Right: s2}
	s8 := &Node{Left: s7, Right: s6}
	t1 := &Tree{Root: s8}
	f50 := &Node{}

	var testCase = []struct {
		in   *Tree
		in2  *Node
		in3  *Node
		want bool
	}{
		{t1, s2, s7, true},
		{t1, f50, s7, false},
	}

	for _, tt := range testCase {
		got := QuestionOne(tt.in, tt.in2, tt.in3)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}
