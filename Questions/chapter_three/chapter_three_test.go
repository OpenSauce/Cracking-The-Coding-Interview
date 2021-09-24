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

func TestQuestionFive(t *testing.T) {
	var s1 Stack
	s1.push(2)
	s1.push(1)
	s1.push(4)
	s1.push(8)
	s1.push(9)
	s1.push(6)
	s1.push(5)
	s1.push(7)
	s1.push(3)
	var s2 Stack
	s2.push(9)
	s2.push(8)
	s2.push(7)
	s2.push(6)
	s2.push(5)
	s2.push(4)
	s2.push(3)
	s2.push(2)
	s2.push(1)

	var testCase = []struct {
		in   Stack
		want bool
	}{
		{s1, true},
	}

	for _, tt := range testCase {
		val := SortStack(tt.in)
		val.Display()
		s2.Display()
		got := val.equals(s2)
		if got != tt.want {
			t.Errorf("got: %v; want: %v", got, tt.want)
		}

	}
}

func TestQuestionSix(t *testing.T) {
	myList := List{}
	myList.Enqueue(Dog{})
	myList.Enqueue(Cat{})
	myList.Enqueue(Dog{})
	myList.Enqueue(Cat{})
	myList.Display()
	myList.DequeueType("cat")
	myList.Display()
	myList.Dequeue()
	myList.Display()
	myList.Enqueue(Dog{})
	myList.Enqueue(Cat{})
}
