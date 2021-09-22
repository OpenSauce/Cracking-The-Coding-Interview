package chapter_three

import (
	"fmt"
)

type Stack struct {
	top      *StackNode
	capacity int
}

type StackNode struct {
	data interface{}
	next *StackNode
}

var g_stack Stack

func (s *Stack) pop() (interface{}, error) {
	if s.top != nil {
		data := s.top.data
		s.top = s.top.next
		return data, nil
	}
	return 0, fmt.Errorf("empty Stack")
}

func (s *Stack) push(in interface{}) {
	t := StackNode{data: in, next: s.top}
	s.top = &t
}

func (s *Stack) peek() (interface{}, error) {
	if s.top != nil {
		return s.top.data, nil
	}
	return 0, fmt.Errorf("empty Stack")
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) equals(s2 Stack) bool {
	st := s.top
	st2 := s2.top
	for st != nil || st2 != nil {
		if st == nil || st2 == nil {
			return false
		}
		if st.data != st2.data {
			return false
		}
		st = st.next
		st2 = st2.next
	}
	return true
}

func (s *Stack) Display() {
	fmt.Println("----")
	st := s.top
	for st != nil {
		fmt.Println(st.data)
		st = st.next
	}
}

func setPop() (interface{}, error) {
	s := g_stack.top.data.(Stack)
	data, err := s.pop()
	if err != nil {
		g_stack.top = s.top.next
		s := g_stack.top.data.(Stack)
		s.capacity--
		g_stack.top.data = s
		return s.pop()
	}
	s.capacity--
	g_stack.top.data = s
	return data, fmt.Errorf("empty Stack")
}

func setPush(in int) {
	if g_stack.isEmpty() {
		newStack := Stack{capacity: 0}
		g_stack.push(newStack)
	}

	s := g_stack.top.data.(Stack)
	if s.capacity == 5 {
		newStack := Stack{capacity: 1}
		newStack.push(in)
		g_stack.push(newStack)
		return
	}
	s.capacity++
	s.push(in)
	g_stack.top.data = s
}

//----------------------------------------------------------------------------------------------------------------------------------

type MyQueue struct {
	StackOne Stack
}

func (q *MyQueue) Add(in interface{}) {
	if q.StackOne.top == nil {
		q.StackOne.push(in)
		return
	}

	st := q.StackOne.top
	var temp Stack
	for st != nil {
		temp.push(st.data)
		st = st.next
		q.StackOne.pop()
	}
	temp.push(in)

	t := temp.top
	for t != nil {
		q.StackOne.push(t.data)
		t = t.next
	}
}

func (q *MyQueue) Remove() {
	if q.StackOne.top == nil {
		return
	}

	q.StackOne.pop()
}

func (q *MyQueue) Display() {
	fmt.Println("----")
	st := q.StackOne.top
	for st != nil {
		fmt.Println(st.data)
		st = st.next
	}
}

func QuestionOne(word string) {
	/*
		Use the modulus and divide by the number of stacks. Track the head of each stack.
	*/
}

func QuestionTwo(word string) {
	/*
		Have another stack which is used to store each minimum, or use an extra datapoint to identify which is the minimum.
	*/
}

func SortStack(s Stack) Stack {
	count := 0
	st := s.top
	for st != nil {
		count++
		st = st.next
	}

	var temp Stack
	stack := s
	for i := 0; i < count; i++ {
		max := 0
		st = stack.top
		for st != nil {
			if st.data.(int) > max {
				max = st.data.(int)
			}
			st = st.next
		}
		if max > 0 {
			temp.push(max)
		}
		st = stack.top
		for st != nil {
			if st.data.(int) == max {
				st.data = 0
			}
			st = st.next
		}
	}
	return temp
}
