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
