package main

import (
	"fmt"
	"layi"
)

func main() {
	// if len(os.Args) <= 1 {
	// 	return
	// }
	// argss := os.Args[1]
	// arggs2 := os.Args[2]
	// layi.Union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	fmt.Println(layi.Print(5))
	layi.TowerOfHanoi(4, 'A', 'B', 'C')
}
