package methods

func Primes(minimum, maximum int) []int {
	primes := make([]int, 0, maximum-minimum+1)

	for i := minimum; i <= maximum; i++ {
		isPrime := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
