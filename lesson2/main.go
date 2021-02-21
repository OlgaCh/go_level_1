package main

import "fmt"

func main() {
	fmt.Println("Выберите вариант программы: [1, 2, 3]")
	fmt.Println("1 - Подсчет площади прямоугольника.")
	fmt.Println("2 - Подсчет диаметра и длины окружности круга.")
	fmt.Println("3 - Разбиение трехзначного числа на сотни, десятки и единицы.")

	var option string
	_, _ = fmt.Scan(&option)

	switch option {
	case "1":
		Rectangle()
	case "2":
		Circle()
	case "3":
		Number()
	default:
		fmt.Println("Неизвестная опция ", option)
	}

}