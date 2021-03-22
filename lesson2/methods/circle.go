package methods

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
	diameter := getDiameter(radius)
	length := getLength(diameter)

	fmt.Println("Диаметр:", diameter)
	fmt.Println("Длина окружности:", length)
}

func getDiameter(radius float64) float64 {
	return 2 * radius
}

func getLength(diameter float64) float64 {
	return diameter * math.Pi
}
