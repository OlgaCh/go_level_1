package main

import (
	"fmt"
	"github.com/OlgaCh/go_level_1/lesson3/prime"
)

func main() {
	iterations := 1000
	fmt.Println("Go Math/Big")
	fmt.Println("Prime numbers from 0 till ", iterations, prime.GoPrimeImpl(iterations))
	fmt.Println("Custom")
	fmt.Println("Prime numbers from 0 till ", iterations, prime.CustomPrimeImpl(iterations))
}
