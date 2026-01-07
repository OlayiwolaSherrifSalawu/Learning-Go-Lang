package stacks

import "fmt"

var maxSize int = 10

type ArrayStack struct {
	top     int
	capcity int
	array   []*int
}

func createStack() *ArrayStack {
	S := &ArrayStack{}
	S.capcity = maxSize
	S.top = -1
	S.array = make([]*int, S.capcity)
	return S
}

func IsEmptyStack(S *ArrayStack) bool {
	return (S.top == -1)
}
func IsFullStack(S *ArrayStack) bool {
	return (S.top == S.capcity-1)
}
func Push(S *ArrayStack, data int) {
	if IsFullStack(S) {
		fmt.Println("Stack OverFlow")
		return
	} else {
		S.array[S.top+1] = &data
	}
}
func Pop(S *ArrayStack) {
	if IsEmptyStack(S) {
		fmt.Println("Stack underFlow")
		return
	} else {
		S.array = S.array[0 : S.top-1]
	}
}
func DeleteStack(S *ArrayStack) {
	S.array = nil
	S = nil
}

// SOlving a problem
// Question : Discuss how stacks can be used for checking balancing of symbols.

// Solution: Stacks can be used to check whether the given expression has balanced symbols. This
// algorithm is very useful in compilers. Each time the parser reads one character at a time. If the
// character is an opening delimiter such as (, {, or [- then it is written to the stack. When a closing
// delimiter is encountered like ), }, or ]-the stack is popped.
