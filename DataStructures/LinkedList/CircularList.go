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
