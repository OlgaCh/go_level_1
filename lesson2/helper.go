package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validateFloatInput() (float64, error) {
	var s string

	_, errScan := fmt.Scan(&s)
	n, err := strconv.ParseFloat(s, 8)
	if n <= 0 {
		err = errors.New("negative number")
	}
	if errScan != nil || err != nil {
		fmt.Println("Введите число больше 0!")
	}
	return n, err
}
