package main

// linked list
type Node struct {
	data int
	next *Node
}
type list struct {
	head *Node
	tail *Node
}

func insertFront(l *list, data int) {
	// current :=
	node := &Node{data: data}
	if node == nil {
		return
	}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	} else {
		node.next = l.head
		l.head = node
	}

}

func insertBack(l *list, data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func main() {

}
