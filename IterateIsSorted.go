package layi

func IterateIsSorted(arr []int) bool {
	for i := 0; i < len(arr); i++ { //O(n)
		for j := i; j < i+1; j++ { //O(logn)
			if j+1 < len(arr) {
				if arr[j+1] < arr[i] {
					return false
				}
			}
		}
	}
	return true
}

// Time Complexity O(n)
