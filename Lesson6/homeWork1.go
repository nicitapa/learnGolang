package Lesson6

import "fmt"

// Сумма двух чисел
func Sum(a, b int) int {
	return a + b
}

// Деление с остатком
func Divide(a, b int) (int, int) {
	return a / b, a % b
}

// Функция как аргумент
func CallFunction(f func()) {
	f()
}

// Массив из 5 целых чисел
func PrintArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, val := range arr {
		fmt.Println(val)
	}
}

// Удвоение элементов среза
func DoubleSlice(slice []int) []int {
	doubled := make([]int, len(slice))
	for i, val := range slice {
		doubled[i] = val * 2
	}
	return doubled
}

// Объединение двух срезов
func MergeSlices(a, b []int) []int {
	return append(a, b...)
}

// Хранение возраста по имени
func PrintAges() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   35,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

// Подсчёт слов
func CountWords(words []string) map[string]int {
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	return count
}

// Удаление ключа и проверка
func RemoveKey(m map[string]int, key string) {
	delete(m, key)
	if _, exists := m[key]; !exists {
		fmt.Printf("Key '%s' was removed\n", key)
	}
}

// Структура Person
type Person struct {
	Name   string
	Salary float64
}

// Метод для отображения информации
func (p Person) PrintInfo() {
	fmt.Printf("Name: %s, Salary: %.2f\n", p.Name, p.Salary)
}
