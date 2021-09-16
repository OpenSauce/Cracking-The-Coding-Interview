package chapter_two

import (
	"strings"
	"testing"
)

func TestQuestionOne(t *testing.T) {
	inputList := createLinkedList(2, 3, 3, 4, 5, 3, 6, 4, 7)
	outputList := createLinkedList(2, 3, 4, 5, 6, 7)

	var testCase = []struct {
		in   List
		want List
	}{
		{inputList, outputList},
	}

	for _, tt := range testCase {
		got := QuestionOne(tt.in)
		if got.String() != tt.want.String() {
			t.Errorf("got: %s; want: %s", got.String(), tt.want.String())
		}
	}
}

func TestQuestionTwo(t *testing.T) {
	inputList := createLinkedList(2, 3, 3, 4, 5, 3, 6, 4, 7)
	inputList2 := createLinkedList(2, 3, 3, 4, 5, 3, 6, 4, 7)

	var testCase = []struct {
		in   List
		in2  int
		want int
	}{
		{inputList, 3, 6},
		{inputList2, 5, 5},
	}

	for _, tt := range testCase {
		got := QuestionTwo(tt.in, tt.in2)
		if got != tt.want {
			t.Errorf("got: %d; want: %d", got, tt.want)
		}
	}
}

func TestQuestionThree(t *testing.T) {
	inputList := createLinkedList(2, 3, 4, 5, 6, 7)
	outputList := createLinkedList(2, 4, 5, 6, 7)

	var testCase = []struct {
		in   List
		in2  int
		want List
	}{
		{inputList, 3, outputList},
	}

	for _, tt := range testCase {
		got := QuestionThree(tt.in, tt.in2)
		if got.String() != tt.want.String() {
			t.Errorf("got: %s; want: %s", got.String(), tt.want.String())
		}
	}
}

func TestQuestionFour(t *testing.T) {
	inputList := createLinkedList(2, 3, 4, 5, 6, 7, 8, 9, 10)

	var testCase = []struct {
		in  List
		in2 int
	}{
		{inputList, 5},
	}

	for _, tt := range testCase {
		got1, got2 := QuestionFour(tt.in, tt.in2)
		if strings.ContainsAny(got1.String(), "56789") {
			t.Errorf("got: %s; want: %s", got1.String(), "56789")
		}
		if strings.ContainsAny(got2.String(), "234") {
			t.Errorf("got: %s; want: %s", got2.String(), "234")
		}
	}
}

func TestQuestionFive(t *testing.T) {
	input1 := createLinkedList(7, 1, 6)
	input2 := createLinkedList(5, 9, 2)
	output := createLinkedList(2, 1, 9)

	var testCase = []struct {
		in   List
		in2  List
		want List
	}{
		{input1, input2, output},
	}

	for _, tt := range testCase {
		got := QuestionFive(tt.in, tt.in2)
		if got.String() != tt.want.String() {
			t.Errorf("got: %s; want: %s", got.String(), tt.want.String())
		}

	}
}

func TestQuestionSix(t *testing.T) {
	input1 := createLinkedList(7, 1, 7)
	input2 := createLinkedList(5, 9, 2)

	var testCase = []struct {
		in   List
		want bool
	}{
		{input1, true},
		{input2, false},
	}

	for _, tt := range testCase {
		got := QuestionSix(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func TestQuestionSeven(t *testing.T) {
	// input1 := createLinkedList(7, 1, 7)
	// input2 := createLinkedList(5, 9, 2)
	// input3 := createLinkedList(5, 9, 2)
	// input4 := createLinkedList(5, 9, 2)

	// var testCase = []struct {
	// 	in   List
	// 	in2  List
	// 	want int
	// }{
	// 	{input1, input2, 5},
	// 	{input3, input4, 0},
	// }

	// for _, tt := range testCase {
	// 	got := QuestionSeven(tt.in, tt.in2)
	// 	if got != tt.want {
	// 		t.Errorf("got: %v; want: %v", got, tt.want)
	// 	}

	// }
}

func TestQuestionEight(t *testing.T) {
	input1 := createLinkedList(7, 1, 7)
	input2 := createLinkedList(5, 9, 2)

	var testCase = []struct {
		in   List
		want int
	}{
		{input1, 7},
		{input2, 0},
	}

	for _, tt := range testCase {
		got := QuestionEight(tt.in)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func createLinkedList(vals ...int) List {
	list := List{}
	for _, v := range vals {
		list.AppendToTail(v)
	}
	return list
}
