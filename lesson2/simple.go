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
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
