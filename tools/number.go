package tools

func IsOdd(n int) bool {
	return n%2 == 1
}

func IsEven(n int) bool {
	return n%2 == 0
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
