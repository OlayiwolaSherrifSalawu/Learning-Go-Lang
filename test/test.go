package main

import (
	"layi"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	argss := os.Args[1]
	arggs2 := os.Args[2]
	layi.Union(argss, arggs2)
}
