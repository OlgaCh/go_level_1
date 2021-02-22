package main

import "fmt"

func main() {
	fmt.Println("Выберите вариант программы: [1, 2, 3, 4, 5]")
	fmt.Println("1 - Подсчет площади прямоугольника.")
	fmt.Println("2 - Подсчет диаметра и длины окружности круга.")
	fmt.Println("3 - Разбиение трехзначного числа на сотни, десятки и единицы.")
  fmt.Println("4 - Калькулятор с 4 операциями (сложение, вычитание, умножение, деление).")
	fmt.Println("5 - Простое число.")

	var option string
	_, _ = fmt.Scan(&option)

	switch option {
	case "1":
		Rectangle()
	case "2":
		Circle()
	case "3":
		Number()
  case "4":
		Calculator()
	case "5":
		PrimeNumber()
	default:
		fmt.Println("Неизвестная опция ", option)
	}

}
