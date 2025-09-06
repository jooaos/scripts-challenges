package main

import "fmt"

type Node struct {
	startTime int
	endTime   int
	left      *Node
	right     *Node
}

func NewNode(startTime int, endTime int) *Node {
	return &Node{
		startTime: startTime,
		endTime:   endTime,
	}
}

type MyCalendar struct {
	head *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (m *MyCalendar) Book(startTime int, endTime int) bool {
	var ok bool
	m.head, ok = m.insert(m.head, startTime, endTime)
	return ok
}

func (m *MyCalendar) insert(node *Node, startTime, endTime int) (*Node, bool) {
	if node == nil {
		return NewNode(startTime, endTime), true
	}

	ok := false
	if endTime <= node.startTime {
		node.left, ok = m.insert(node.left, startTime, endTime)
	} else if startTime >= node.endTime {
		node.right, ok = m.insert(node.right, startTime, endTime)
	}

	return node, ok
}

func main() {
	obj := Constructor()
	param_1 := obj.Book(10, 20)
	fmt.Println(param_1)
	param_2 := obj.Book(15, 25)
	fmt.Println(param_2)
	param_3 := obj.Book(20, 30)
	fmt.Println(param_3)
}
