package main

import (
	"fmt"
)

func Rectangle() {
	fmt.Println("Введите сторону А:")
	a, err := validateFloatInput()
	if err != nil {
		return
	}

	fmt.Println("Введите сторону B:")
	b, err := validateFloatInput()
	if err != nil {
		return
	}

	fmt.Println("Площадь прямоугольника:", a*b)
}
