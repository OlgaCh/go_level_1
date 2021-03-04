package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func FibonacciWithDefer() func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a = a + b
			a, b = b, a
		}()
		return a
	}
}

func FibonacciWithCache(n int, cache map[int]int) int {
	if val, found := cache[n]; found {
		return val
	}
	cache[n] = FibonacciWithCache(n-1, cache) + FibonacciWithCache(n-2, cache)
	return cache[n]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	var s string

	fmt.Println("Введите порядковый номер числа из ряда Фибоначчи.")
	_, errScan := fmt.Scan(&s)
	n, err := strconv.ParseInt(s, 10, 0)
	if n <= 0 || errScan != nil || err != nil {
		fmt.Println("Число должно быть целым и положительным!")
		return
	}

	func() {
		defer timeTrack(time.Now(), "FibonacciWithDefer")
		f := FibonacciWithDefer()
		for i := 0; i < int(n)-1; i++ {
			f()
		}

		fmt.Println("Число Фибоначчи:", f())
	}()

	func() {
		defer timeTrack(time.Now(), "FibonacciWithCache")
		fmt.Println("Число Фибоначчи:", FibonacciWithCache(int(n), map[int]int{1: 1, 2: 1}))
	}()
}
