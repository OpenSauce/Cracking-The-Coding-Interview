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
