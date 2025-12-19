package main

import (
	"fmt"
	"layi"
)

func main() {
	arr := []int{0, 1, 3, 9, 4, 5, 6, 7, 8, 9}
	testIsSorted := layi.IterateIsSorted(arr)
	fmt.Println(testIsSorted)
}
