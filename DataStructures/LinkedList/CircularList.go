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
	for current.next != l.head {
		count++
		current = current.next
	}
	fmt.Println(count)
}
func insertCf(l *Clist, data int) {

}

func main() {
	node := &Clist{}
	if node.head == nil {
		return
	}

	CountNode(node)

}
