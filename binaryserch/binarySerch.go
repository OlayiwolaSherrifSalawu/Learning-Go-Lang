package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valid, num := RecuBinarySearch(arr, 9)
	fmt.Println(valid, num)
}

func RecuBinarySearch(arr []int, value int) (bool, int) {
	if len(arr) == 0 {
		return false, 0
	}
	lenArr := len(arr)
	if value < arr[len(arr)/2] {
		return RecuBinarySearch(arr[:len(arr)/2], value)
	} else if value > arr[len(arr)/2] {
		return RecuBinarySearch(arr[lenArr/2:], value)
	} else if value == arr[len(arr)/2] {
		return true, value
	} else {
		return false, value
	}
}
