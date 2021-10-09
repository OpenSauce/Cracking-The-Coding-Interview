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

func TestQuestionTwo(t *testing.T) {
	input := []int{1, 2, 4, 5, 6, 7, 8}

	bst := QuestionTwo(input)

	search(bst.Root)
}

func TestQuestionThree(t *testing.T) {
	s1 := &Node{}
	s2 := &Node{}
	s3 := &Node{Left: s1}
	s4 := &Node{}
	s5 := &Node{Left: s4}
	s6 := &Node{Left: s3}
	s7 := &Node{Left: s5, Right: s2}
	s8 := &Node{Left: s7, Right: s6}
	t1 := &Tree{Root: s8}

	QuestionThree(t1)
}

func TestQuestionFour(t *testing.T) {
	s1 := &Node{}
	s2 := &Node{Left: s1}
	s3 := &Node{Left: s2}
	s4 := &Node{}
	s5 := &Node{Left: s4}
	s6 := &Node{Left: s3}
	s7 := &Node{Left: s5, Right: s2}
	s8 := &Node{Left: s7, Right: s6}
	t1 := &Tree{Root: s8}

	n2 := &Node{}
	n3 := &Node{}
	n4 := &Node{}
	n5 := &Node{}
	n6 := &Node{Left: n3, Right: n2}
	n7 := &Node{Left: n5, Right: n4}
	n8 := &Node{Left: n7, Right: n6}
	t2 := &Tree{Root: n8}

	var testCase = []struct {
		in   *Node
		want bool
	}{
		{t1.Root, false},
		{t2.Root, true},
	}

	for _, tt := range testCase {
		got := QuestionFour(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func TestQuestionFive(t *testing.T) {
	s1 := &Node{}
	s2 := &Node{}
	s3 := &Node{Left: s1}
	s4 := &Node{}
	s5 := &Node{Left: s4}
	s6 := &Node{Left: s3}
	s7 := &Node{Left: s5, Right: s2, Data: 10}
	s8 := &Node{Left: s7, Right: s6, Data: 5}
	t1 := &Tree{Root: s8}

	n6 := &Node{Data: 5}
	n7 := &Node{Data: 2}
	n8 := &Node{Left: n7, Right: n6, Data: 3}
	t2 := &Tree{Root: n8}

	var testCase = []struct {
		in   *Node
		want bool
	}{
		{t1.Root, false},
		{t2.Root, true},
	}

	for _, tt := range testCase {
		got := QuestionFive(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func TestQuestionSix(t *testing.T) {
	s1 := &Node{Data: 1}
	s2 := &Node{Data: 12}
	s3 := &Node{Left: s1, Data: 3}
	s4 := &Node{Data: 4}
	s5 := &Node{Left: s4, Data: 5}
	s6 := &Node{Left: s3, Data: 6}
	s7 := &Node{Left: s5, Right: s2, Data: 7}
	s8 := &Node{Left: s7, Right: s6, Data: 8}
	t1 := &Tree{Root: s8}

	var testCase = []struct {
		in   *Node
		want int
	}{
		{t1.Root, 1},
	}

	for _, tt := range testCase {
		got := QuestionSix(tt.in)
		if got.Data != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func TestQuestionSeven(t *testing.T) {
	// todo Implement test for the project graph
}

func TestQuestionEight(t *testing.T) {

}
