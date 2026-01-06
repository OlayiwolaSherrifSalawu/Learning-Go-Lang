package stacks

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}
type StackList struct {
	head *StackNode
}

func PshLinkedStack(S *StackList, data int) {
	node := &StackNode{data: data}
	if S.head == nil {
		S.head = node
		return
	}
	node.next = S.head
	S.head = node
}
func PopLinkedStack(S *StackList) {
	if S.head == nil {
		fmt.Println("Stack Underflow")
	}
	temp := S.head.next
	S.head = temp
	// S.head.next = nil
	// S.head = nil
}
