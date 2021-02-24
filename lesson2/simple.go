package main

import (
	"fmt"
	"strconv"
)

func PrimeNumber() {
	fmt.Println("Введите целое число:")
	var s string

	_, errScan := fmt.Scan(&s)
	n, err := strconv.ParseInt(s, 10, 64)
	if errScan != nil || err != nil {
		fmt.Println("Ошибка при вводе числа")
		return
	}
	for i := 0; i <= int(n); i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
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
