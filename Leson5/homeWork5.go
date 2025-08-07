package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка: деление на ноль невозможно")
	}
	return a / b, nil
}

func main() {
	for {
		var a, b int
		fmt.Println("Введите два целых числа (a и b):")

		// Проверка корректности ввода
		_, err := fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите два целых числа.")
			continue
		}

		result, err := divide(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Результат деления: %d\n", result)

			// Выбор сообщения в зависимости от результата
			switch {
			case result > 10:
				fmt.Println("Результат большой")
			case result >= 1 && result <= 10:
				fmt.Println("Результат средний")
			default:
				fmt.Println("Результат маленький или ноль")
			}
		}

		// Повтор операции
		var answer string
		fmt.Println("Хотите повторить операцию? (да/нет):")
		fmt.Scanln(&answer)
		if answer != "да" {
			break
		}
	}
}
