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
	current := l.head
	for current != l.head {
		count++
		current = current.next
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

func main() {
	node := &Clist{}
	insertCf(node, 1)
	insertCf(node, 2)
	insertCf(node, 3)
	insertCf(node, 4)
	insertCf(node, 5)
	insertCf(node, 6)
	insertCf(node, 7)
	insertCf(node, 8)

	CountNode(node)
	current := node.head
	for current.next != node.head {
		fmt.Print(current.data)
		fmt.Print("->")
		current = current.next
	}
	fmt.Println()
}
