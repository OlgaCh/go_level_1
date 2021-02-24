package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
)

func main() {
	fmt.Println("Go Math/Big")
	goPrimeImpl()
	fmt.Println("Custom")
	customPrimeImpl()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func goPrimeImpl() {
	defer timeTrack(time.Now(), "Go Prime")
	totalPrimeNumbers := 0
	for i := 0; i <= int(math.Pow(10, 6)); i++ {
		if big.NewInt(int64(i)).ProbablyPrime(10) {
			totalPrimeNumbers++
		}
	}
	fmt.Println("Prime numbers from 0 to 10^6", totalPrimeNumbers)
}

func customPrimeImpl() {
	defer timeTrack(time.Now(), "Custom Prime")
	totalPrimeNumbers := 0
	for i := 0; i <= int(math.Pow(10, 6)); i++ {
		if isPrime(i) {
			totalPrimeNumbers++
		}
	}
	fmt.Println("Prime numbers from 0 to 10^6", totalPrimeNumbers)
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
