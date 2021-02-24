package main

import (
	"fmt"
	"math"
)

func Circle() {
	fmt.Println("Введите площадь круга:")
	a, err := validateFloatInput()
	if err != nil {
		return
	}

	radius := math.Sqrt(a / math.Pi)
	diameter := 2 * radius
	length := diameter * math.Pi

	fmt.Println("Диаметр:", diameter)
	fmt.Println("Длина окружности:", length)
}
