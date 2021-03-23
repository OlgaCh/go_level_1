package prime

import (
	"math/big"
)

func GoPrimeImpl(iterations int) int {
	totalPrimeNumbers := 0
	for i := 0; i <= iterations; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(10) {
			totalPrimeNumbers++
		}
	}
	return totalPrimeNumbers
}

func CustomPrimeImpl(iterations int) int {
	totalPrimeNumbers := 0
	for i := 0; i <= iterations; i++ {
		if isPrime(i) {
			totalPrimeNumbers++
		}
	}
	return totalPrimeNumbers
}

func isPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
