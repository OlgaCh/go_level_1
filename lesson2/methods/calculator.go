package methods

import (
	"fmt"
	"strconv"
)

func Calculator() {
	fmt.Println("Введите первое число:")
	a, err := validateCalcInput()
	if err != nil {
		return
	}

	fmt.Println("Введите второе число:")
	b, err := validateCalcInput()
	if err != nil {
		return
	}

	fmt.Println("Выберите допустимую операцию: +, -, *, /")
	var operation string
	_, _ = fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Деление на ноль не допустимо!")
		}
	default:
		fmt.Println("Неизвестная операция ", operation)
	}
}

func validateCalcInput() (float64, error) {
	var s string

	_, inputErr := fmt.Scan(&s)
	n, err := strconv.ParseFloat(s, 8)
	if inputErr != nil || err != nil {
		fmt.Println("Введите число!")
	}
	return n, err
}
