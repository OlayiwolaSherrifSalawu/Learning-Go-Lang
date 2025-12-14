package layi

import "strconv"

func Fprime(nb int) string {
	value := nb
	result := ""
	for i := 2; i*i <= nb; i++ {
		for value%i == 0 && value > 0 {
			result += strconv.Itoa(i) + "*"
			value = value / i
		}
	}
	result = result[:len(result)-1]
	return result
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}
