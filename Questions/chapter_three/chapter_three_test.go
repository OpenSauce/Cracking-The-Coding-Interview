package chapter_three

import (
	"fmt"
	"testing"
)

func TestQuestionOne(t *testing.T) {
	//Not a function question
}

func TestQuestionTwo(t *testing.T) {
	//Not a function question
}

func TestQuestionThree(t *testing.T) {
	setPush(1)
	setPush(2)
	setPush(3)
	setPush(4)
	setPush(5)
	setPop()
	setPush(1)
	fmt.Printf("%+v\n", g_stack.top)
	setPush(2)
	fmt.Printf("%+v\n", g_stack.top)
}

func TestQuestionFour(t *testing.T) {
	var q MyQueue
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)
	q.Add(6)
	q.Add(7)
	q.Display()
	q.Remove()
	q.Add(8)
	q.Display()
}
