package main

import "fmt"

type Cll struct {
	data int
	next *Cll
}

type Clist struct {
	head *Cll
}

func CountNode(l *Clist) {
	count := 0
	if l.head == nil {
		fmt.Println(count)
		return
	}
	// prev := l.head.next
	current := l.head
	for {
		count++
		current = current.next
		if current == l.head {
			break
		}
	}
	fmt.Println(count)
}
func insertCf(l *Clist, data int) {
	node := &Cll{data: data}
	if l.head == nil {
		l.head = node
	}
	node.next = l.head
	thehead := l.head
	for thehead.next != l.head {
		node.next = node
		thehead = thehead.next
	}
	node.next = l.head
	thehead.next = node

}
func circularFront(l *Clist, data int) {
	node := &Cll{data: data}
	if l.head == nil {
		l.head = node
	}
	node.next = l.head
	tailNode := l.head
	for tailNode.next != l.head {
		tailNode = tailNode.next
	}
	tailNode.next = node
	l.head = node
}
func deletelastNode(l Clist) {
	prev := l.head
	tailNode := l.head.next
	for tailNode.next != l.head {
		prev = prev.next
		tailNode = tailNode.next
	}
	prev.next = l.head
	tailNode.next = nil
	tailNode = nil
}
func deleteFirstNode(l Clist) {
	if l.head == nil {
		return
	}
	// temp := l.head
	if l.head.next == l.head {
		l.head = nil
	}
	tailNode := l.head
	for tailNode.next != l.head {
		tailNode = tailNode.next
	}
	l.head = l.head.next
	tailNode.next = l.head

	// temp = nil
	// temp.next = nil
}
func main() {
	node := &Clist{}
	circularFront(node, 1)
	circularFront(node, 2)
	circularFront(node, 3)
	circularFront(node, 4)
	circularFront(node, 5)
	circularFront(node, 6)
	circularFront(node, 7)
	circularFront(node, 8)
	deleteFirstNode(*node)
	CountNode(node)
	current := node.head
	for {
		fmt.Print(current.data)
		current = current.next
		fmt.Print("->")
		if current == node.head {
			break
		}
	}
	fmt.Println()
}
