package fprime

func Fprime(nb int) bool {
	if IsPrime(nb) {
		return true
	}
	return false
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
