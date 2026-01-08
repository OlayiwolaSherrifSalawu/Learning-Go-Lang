package main

import "fmt"

func main() {
	// if len(os.Args) <= 1 {
	// 	return
	// }
	// argss := os.Args[1]
	// arggs2 := os.Args[2]
	// layi.Union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	// fmt.Println(layi.Print(5))
	// layi.TowerOfHanoi(4, 'A', 'B', 'C')
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// testIsSorted := layi.IterateIsSorted(arr)
	// fmt.Println(testIsSorted)
	var items map[byte]byte = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	val, ok := items[')']
	fmt.Println(val, ok)
}
