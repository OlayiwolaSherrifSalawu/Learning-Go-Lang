package main

import "fmt"

var x int = 5

func add(arg ...int) int {
	total := 0
	for _, value := range arg {
		total += value
	}
	return total
}

//

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func halveEven(x int) {
	half := x / 2
	if half%2 == 0 {
		fmt.Println(half, "True")
	}
	fmt.Println(half, "False")

}

func highestNumber(arg ...int) int {
	highest := arg[0]

	for _, num := range arg {
		if num > highest {
			highest = num
		}
	}
	return highest
}
func fbl(n int) int {
	if n == 0 {
		return 0
	}
	return fbl(n-1) + fbl(n-2)
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	halveEven(23)
	fmt.Println(highestNumber(1, 2, 3, 4, 189999, 23, 243, 344, 8, 3333))
}
