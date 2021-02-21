package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validateFloatInput() (float64, error) {
	var s string
	var n float64

	_, err := fmt.Scan(&s)
	n, err = strconv.ParseFloat(s, 8)
	if n <= 0 {
		err = errors.New("negative number")
	}
	if err != nil {
		fmt.Println("Введите число больше 0!")
	}
	return n, err
}
