package main

import "fmt"

var Array [3]int

func main() {
	KString(3, 3)
	fmt.Println()
}

func KString(n, k int) {
	if n < 1 {
		fmt.Print(Array)
	} else {
		for i := 0; i < k; i++ {
			Array[n-1] = i
			KString(n-1, k)
		}
	}
}
