package methods

import (
	"errors"
	"fmt"
	"strconv"
)

func Number() {
	fmt.Println("Введите число от 0 до 999:")
	n, err := validateInput()
	if err != nil {
		return
	}

	hundreds := n / 100 >> 0
	dozens := (n / 10 >> 0) % 10
	units := n % 10

	printNumberParts(hundreds, dozens, units)
}

func validateInput() (int, error) {
	var s string

	_, errScan := fmt.Scan(&s)
	n, err := strconv.ParseInt(s, 10, 64)
	if n < 0 || n > 999 {
		err = errors.New("wrong number")
	}
	if errScan != nil || err != nil {
		fmt.Println("Число должно быть от 0 до 999!")
	}
	return int(n), err
}

func printNumberParts(hundreds int, dozens int, units int) {
	labels := [3]string{"Сотен", "Десятков", "Единиц"}
	parts := [3]int{hundreds, dozens, units}

	for i, s := range parts {
		if s > 0 {
			fmt.Println(labels[i], s)
		}
	}
}
