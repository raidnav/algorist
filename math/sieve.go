package math

func IsPrime(n int) bool {
	primes, flag := func(N int) ([]int, [10_000_000]bool) {
		var isPrime [10_000_000]bool
		var primes []int
		for i := 2; i < N; i++ {
			isPrime[i] = true
		}
		for i := 2; i < N; i++ {
			if isPrime[i] {
				for j := i * i; j < N; j += i {
					isPrime[j] = false
				}
				primes = append(primes, i)
			}
		}
		return primes, isPrime
	}(10_000_000)
	if n < 10_000_000 {
		return flag[n]
	}
	for i := 0; i < len(primes); i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}
