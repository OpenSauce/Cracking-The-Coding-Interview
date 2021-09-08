package chapter_two

import "fmt"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
}

func (l *List) AppendToTail(val int) {
	node := &Node{
		next: nil,
		data: val,
	}

	if l.head == nil {
		l.head = node
		return
	}

	list := l.head
	for list.next != nil {
		list = list.next
	}
	list.next = node
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v -> ", list.data)
		list = list.next
	}
	fmt.Println()
}

func (l *List) String() string {
	var res string
	list := l.head
	for list != nil {
		res += fmt.Sprintf("%+v", list.data)
		list = list.next
	}
	return res
}

func QuestionOne(l List) List {
	previousValues := make(map[int]struct{})
	n := l.head
	for n.next != nil {
		previousValues[n.data] = struct{}{}
		if _, ok := previousValues[n.next.data]; ok {
			if n.next.next != nil {
				n.next = n.next.next
				continue
			}
			n.next = nil
			return l
		}
		n = n.next
	}
	return l
}

func QuestionTwo(l List, k int) int {
	length := 0
	n := l.head
	kth := l.head
	for n.next != nil {
		length++
		if length >= k {
			kth = kth.next
		}
		n = n.next
	}
	return kth.data
}

func QuestionThree(l List, v int) List {
	n := l.head
	for n.next != nil {
		if n.next.data == v {
			if n.next.next != nil {
				n.next = n.next.next
				return l
			}
			n.next = nil
			return l
		}
		n = n.next
	}
	return l
}

func QuestionFour(l List, v int) (List, List) {
	l1 := List{}
	l2 := List{}
	n := l.head
	for n.next != nil {
		if n.data >= v {
			l2.AppendToTail(n.data)
		} else {
			l1.AppendToTail(n.data)
		}
		n = n.next
	}
	return l1, l2
}

func QuestionFive(l1, l2 List) List {
	return l1
}
