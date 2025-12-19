package layi

func RecuIsSorted(arr []int, n int) (bool, int) {
	if n == 1 {
		return true, 1 //O(1)
	}
	if arr[n-1] < arr[n-2] {
		return false, 0 //O(1)
	} else {
		return RecuIsSorted(arr, n-1) //O(n)
	}
}

// Time Complexity O(n)
// the recursive method for is sorted and iteratetive method have the same time complexity
