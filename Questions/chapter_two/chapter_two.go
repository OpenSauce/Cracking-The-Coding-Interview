package chapter_two

import (
	"fmt"
	"strconv"
	"strings"
)

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
	n := l1.head
	m := l2.head
	res := List{}
	remainder := 0
	for n != nil && m != nil {
		val := n.data + m.data + remainder
		if val >= 10 {
			res.AppendToTail(val - 10)
			remainder = val / 10
		} else {
			res.AppendToTail(val)
			remainder = 0
		}
		n = n.next
		m = m.next
	}
	return res
}

func QuestionSix(l List) bool {
	n := l.head
	var st1, st2 string
	for n != nil {
		st1 += strconv.Itoa(n.data)
		temp := st2
		st2 = strconv.Itoa(n.data) + temp
		n = n.next
	}
	return strings.Compare(st1, st2) == 0
}
