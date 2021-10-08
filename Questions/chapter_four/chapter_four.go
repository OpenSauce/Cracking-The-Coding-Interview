package chapter_four

import (
	"fmt"
	"math"
)

type Node struct {
	Data          int
	visited       bool
	marked        bool
	InorderParent *Node
	Right         *Node
	Left          *Node
}

type Tree struct {
	Root *Node
}

//----------------------------------

type ListNode struct {
	prev *ListNode
	next *ListNode
	key  *Node
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (L *List) Insert(key *Node) {
	list := &ListNode{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

//----------------------------------

func (l *List) Enqueue(val *Node) {
	node := &ListNode{
		next: nil,
		key:  val,
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

func (l *List) Dequeue() *Node {
	list := l.head
	l.head = list.next
	return list.key
}

func (l *List) isEmpty() bool {
	return l.head == nil
}

//----------------------------------

func QuestionOne(t1 *Tree, n1, n2 *Node) bool {
	search(t1.Root)

	return n1.visited && n2.visited
}

func search(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	root.visited = true
	if root.Left != nil && !root.Left.visited {
		search(root.Left)
	}
	if root.Right != nil && !root.Right.visited {
		search(root.Right)
	}
}

func QuestionTwo(input []int) *Tree {
	myTree := &Tree{}
	myTree.Root = InsertNode(input, 0, len(input)-1)
	return myTree
}

func InsertNode(input []int, start, end int) *Node {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	n := &Node{Data: input[mid]}
	n.Left = InsertNode(input, start, mid-1)
	n.Right = InsertNode(input, mid+1, end)
	return n
}

func QuestionThree(t1 *Tree) {
	BST(t1.Root)
}

func BST(root *Node) {
	queue := List{}
	queue.Enqueue(root)
	current := List{}
	if root != nil {
		current.Insert(root)
	}
	root.marked = true

	for !queue.isEmpty() {
		current.Display()
		r := queue.Dequeue()
		r.visited = true
		current = List{}
		if r.Left != nil && !r.Left.marked {
			r.Left.marked = true
			queue.Enqueue(r.Left)
			current.Insert(r.Left)
		}
		if r.Right != nil && !r.Right.marked {
			r.Right.marked = true
			queue.Enqueue(r.Right)
			current.Insert(r.Right)
		}
	}
}

func QuestionFour(root *Node) bool {
	return checkHeight(root) != math.MinInt
}

func checkHeight(root *Node) int {
	if root == nil {
		return -1
	}

	leftHeight := checkHeight(root.Left)
	if leftHeight == math.MinInt {
		return math.MinInt
	}

	rightHeight := checkHeight(root.Right)
	if rightHeight == math.MinInt {
		return math.MinInt
	}

	heightDiff := leftHeight - rightHeight
	if int(math.Abs(float64(heightDiff))) > 1 {
		return math.MinInt
	} else {
		return int(math.Max(float64(leftHeight), float64(rightHeight)+1))
	}
}

func QuestionFive(root *Node) bool {
	return checkBinarySearch(root)
}

func checkBinarySearch(root *Node) bool {
	if root == nil {
		return true
	}

	if !checkBinarySearch(root.Left) {
		return false
	}
	if !checkBinarySearch(root.Right) {
		return false
	}

	return (root.Left == nil || root.Left.Data < root.Data) && (root.Right == nil || root.Right.Data > root.Data)
}

var previousNode *Node

func QuestionSix(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		return leftMostChild(n.Right)
	} else {
		q := n
		x := q.InorderParent

		for x != nil && x.Left != q {
			q = x
			x = x.InorderParent
		}
		return x
	}
}

func leftMostChild(n *Node) *Node {
	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}
	return n
}

// ------------------------------

type Stack struct {
	top      *StackNode
	capacity int
}

type StackNode struct {
	data interface{}
	next *StackNode
}

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

// ------------------------------

type Graph struct {
	nodes   []Project
	hashMap map[string]Project
}

func (g *Graph) addEdge(startname, endName string) {
	start := g.getOrCreateNode(startname)
	end := g.getOrCreateNode(endName)
	start.addNeighbour(end)
}

func (g *Graph) getOrCreateNode(name string) Project {
	_, ok := g.hashMap[name]
	if !ok {
		node := Project{name: name, hashMap: make(map[string]Project)}
		g.nodes = append(g.nodes, node)
		g.hashMap[name] = node
	}

	return g.hashMap[name]
}

func buildGraph(projects []string, dependencies [][]string) Graph {
	graph := Graph{hashMap: make(map[string]Project)}
	for _, dependency := range dependencies {
		first := dependency[0]
		second := dependency[1]
		graph.addEdge(first, second)
	}

	return graph
}

type Project struct {
	children []Project
	hashMap  map[string]Project
	name     string
	State
}

type State int

const (
	BLANK State = iota
	PARTIAL
	COMPLETE
)

func (p *Project) addNeighbour(node Project) {
	_, ok := p.hashMap[node.name]
	if !ok {
		p.children = append(p.children, node)
		p.hashMap[node.name] = node
	}
}

func (p *Project) getState() State {
	return p.State
}

func (p *Project) setState(s State) {
	p.State = s
}

func (g *Graph) getNodes() []Project {
	return g.nodes
}

func orderProjects(projects []Project) *Stack {
	stack := &Stack{}
	for _, project := range projects {
		if project.getState() == BLANK {
			if !doDFS(project, stack) {
				return nil
			}
		}
	}
	return stack
}

func doDFS(project Project, stack *Stack) bool {
	if project.getState() == PARTIAL {
		return false
	}

	if project.getState() == BLANK {
		project.setState(PARTIAL)
		for _, child := range project.children {
			if !doDFS(child, stack) {
				return false
			}
		}
		project.setState(COMPLETE)
		stack.push(project)
	}
	return true
}

func QuestionSeven(projects []string, dependencies [][]string) *Stack {
	graph := buildGraph(projects, dependencies)
	return orderProjects(graph.getNodes())
}
