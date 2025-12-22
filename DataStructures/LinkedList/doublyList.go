package main

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
	}
	node.next = l.head
	node.prev = nil
	l.head.prev = node
	l.head = node
}
