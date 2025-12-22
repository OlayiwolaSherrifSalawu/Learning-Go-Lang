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

// list insertions
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
	// node.data = data
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}
func insertPosition(l *list, position, data int) {
	k := 1
	node := &Node{data: data}
	current := l.head

	for current != nil && k < position-1 {
		k++
		current = current.next
	}
	node.next = current.next
	current.next = node
}

// list deletion
func deleteFirst(l *list) {
	// node := &Node{}
	temp := l.head.next
	l.head = temp
}
func deletelast(l *list) {
	prev := l.head
	current := l.head.next
	for current != l.tail {
		prev = prev.next
		current = current.next
	}
	prev.next = nil
	l.tail = prev
}
func deletePosition(l *list, position int) {
	k := 1
	prev := l.head
	current := l.head.next
	for prev != nil && k < position-1 {
		k++
		current = current.next
		prev = prev.next
	}
	place := current.next
	prev.next = place
}
func deleteList(l *list) {
	current := l.head
	temp := l.head

	for current != nil {
		l.head = temp
		temp = temp.next
		current = current.next
	}
	l.head = nil
}

// func main() {
// 	listHead := list{}
// 	insertBack(&listHead, 2)
// 	insertBack(&listHead, 3)
// 	insertBack(&listHead, 4)
// 	insertBack(&listHead, 5)
// 	insertBack(&listHead, 6)
// 	insertBack(&listHead, 7)

// 	insertPosition(&listHead, 4, 9)
// 	// deleteFirst(&listHead)
// 	// deletePosition(&listHead, 4)
// 	deleteList(&listHead)

// 	for listHead.head != nil {
// 		data := listHead.head.data
// 		fmt.Print(data)
// 		fmt.Print("->")
// 		listHead.head = listHead.head.next
// 	}
// 	fmt.Println()
// }
