package main

import "fmt"

var A [2]int

func main() {

	nBit(2)
	fmt.Println()
}

func nBit(n int) {
	if n < 1 {
		fmt.Print(A)
	} else {
		A[n-1] = 0
		nBit(n - 1)
		A[n-1] = 1
		nBit(n - 1)
	}

}
