package stacks

import (
	"fmt"
)

var maxSize int = 10

type ArrayStack struct {
	top     int
	capcity int
	array   []*interface{}
}

func createStack() *ArrayStack {
	S := &ArrayStack{}
	S.capcity = maxSize
	S.top = -1
	S.array = make([]*interface{}, S.capcity)
	return S
}

func IsEmptyStack(S *ArrayStack) bool {
	return (S.top == -1)
}
func IsFullStack(S *ArrayStack) bool {
	return (S.top == S.capcity-1)
}
func Push(S *ArrayStack, data interface{}) {
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

/*The opening and closing delimiters are then compared. If they match, the parsing of the string
continues. If they do not match, the parser indicates that there is an error on the line. A linear-time
O(n) algorithm based on stack can be given as:
Algorithm:
a) Create a stack.
b) while (end of input is not reached) {
1) If the character read is not a symbol to be balanced, ignore it.
2) If the character is an opening symbol like (, [, {, push it onto the stack
3) If it is a closing symbol like ),],}, then if the stack is empty report an
error. Otherwise pop the stack.
4) If the symbol popped is not the corresponding opening symbol, report an
error.
}
c) At end of input, if the stack is not empty report an error*/

func balance(Input string) {
	lenGt := len(Input)
	stack := createStack()
	stack.capcity = lenGt
	for i := 0; i < lenGt; i++ {
		if Input[i] == '(' || Input[i] == '[' || Input[i] == '{' {
			Push(stack, Input[i])
		}
		if i != 0 && i-1 >= 0 {
			if Input[i] == ')' && Input[i-1] == '(' {
				PopStack(stack)
			} else if Input[i] == ']' && Input[i-1] == '[' {
				PopStack(stack)
			} else if Input[i] == '}' && Input[i-1] == '{' {
				PopStack(stack)
			} else {
				fmt.Println("error")
			}
		}
	}
	if !IsEmptyStack(stack) {
		fmt.Println("Error")
	}
}
