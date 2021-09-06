package chapter_two

import (
	"testing"
)

func TestQuestionOne(t *testing.T) {
	inputList := List{}
	inputList.AppendToTail(2)
	inputList.AppendToTail(3)
	inputList.AppendToTail(3)
	inputList.AppendToTail(4)
	inputList.AppendToTail(5)
	inputList.AppendToTail(3)
	inputList.AppendToTail(6)
	inputList.AppendToTail(4)
	inputList.AppendToTail(7)

	outputList := List{}
	outputList.AppendToTail(2)
	outputList.AppendToTail(3)
	outputList.AppendToTail(4)
	outputList.AppendToTail(5)
	outputList.AppendToTail(6)
	outputList.AppendToTail(7)

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
