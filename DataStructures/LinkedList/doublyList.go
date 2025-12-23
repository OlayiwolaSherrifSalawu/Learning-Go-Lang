package main

import "fmt"

type DLLNode struct {
	data int
	next *DLLNode
	prev *DLLNode
}
type Dlist struct {
	head *DLLNode
}

// Insertion
func insertDDF(l *Dlist, data int) {
	node := &DLLNode{data: data}
	if l.head == nil {
		l.head = node
		return
	}
	node.next = l.head
	node.prev = nil
	l.head.prev = node
	l.head = node
}
func insertDDB(l *Dlist, data int) {
	node := &DLLNode{data: data}
	if l.head == nil {
		l.head = node
		return
	}
	// prev:=l.head
	current := l.head
	for current.next != nil {
		current = current.next
		// l.head = l.head.next
	}
	// fmt.Println(current.prev)
	node.next = nil
	node.prev = current
	current.next = node
}
func main() {
	node := &Dlist{}
	// ddlist := node.head
	insertDDF(node, 2)
	insertDDF(node, 3)
	insertDDF(node, 4)
	insertDDF(node, 5)
	insertDDF(node, 6)
	for node.head != nil {
		fmt.Print(node.head.data, "->")
		node.head = node.head.next
	}
	fmt.Println()
}
