package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Введите числа для сортировки. После окончания ввода нажмите Ctrl+D")

	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}

	sortArray(inputNums)

	fmt.Println("Отсортированные числа:")
	for _, n := range inputNums {
		fmt.Println(n)
	}

}

func sortArray(arr []int64) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
