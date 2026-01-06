package stacks

import "fmt"

// import
func DoubleStstack(S *ArrayStack) *ArrayStack {
	S.capcity = S.capcity * 2
	return S

}
func IsFullStacks(S *ArrayStack) bool {
	return S.top == S.capcity-1
}

//	func IsEmptyStacks(S *ArrayStack) bool {
//		return S.top == -1
//	}
func PushDynStack(S *ArrayStack, data int) {
	if IsFullStack(S) {
		DoubleStstack(S)
	} else {
		S.array[S.top+1] = &data
	}
}
func PopStack(S *ArrayStack) {
	if IsEmptyStack(S) {
		fmt.Println("stack underflow")
		return
	} else {
		S.array = S.array[0 : S.top-1]
	}
}
