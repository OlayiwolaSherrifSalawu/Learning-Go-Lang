package layi

import "fmt"

func Print(n int) int {
	if n == 0 {
		return 0
	} else {
		fmt.Printf("%d", n)
		return Print(n - 1)
	}

}
