package Lesson6

import (
	"fmt"
)

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

func (e Employee) Info() string {
	return "Имя: " + e.Name + ", Возраст: " + fmt.Sprint(e.Age) +
		", Должность: " + e.Position + ", Зарплата: " + fmt.Sprintf("%.2f", e.Salary)
}

func AddEmployee(list []Employee, newEmployee Employee) []Employee {
	list = append(list, newEmployee)
	return list
}

func FindByName(list []Employee, name string) (Employee, bool) {
	for _, emp := range list {
		if emp.Name == name {
			return emp, true
		}
	}
	return Employee{}, false
}

func GetAverageAge(list []Employee) float64 {
	total := 0
	for _, emp := range list {
		total += emp.Age
	}
	return float64(total) / float64(len(list))
}

func GetAverageSalaries(list []Employee) map[string]float64 {
	sums := make(map[string]float64)
	counts := make(map[string]int)

	for _, emp := range list {
		sums[emp.Position] += emp.Salary
		counts[emp.Position]++
	}

	averages := make(map[string]float64)
	for position := range sums {
		averages[position] = sums[position] / float64(counts[position])
	}

	return averages
}

func main() {
	emp1 := Employee{"Алиса", 30, "Разработчик", 4000}
	emp2 := Employee{"Боб", 25, "Дизайнер", 3500}
	emp3 := Employee{"Карина", 40, "Менеджер", 6000}
	emp4 := Employee{"Дима", 28, "Разработчик", 4200}

	var employees []Employee
	employees = AddEmployee(employees, emp1)
	employees = AddEmployee(employees, emp2)
	employees = AddEmployee(employees, emp3)
	employees = AddEmployee(employees, emp4)

	fmt.Println("Список сотрудников:")
	for _, emp := range employees {
		fmt.Println(emp.Info())
	}

	fmt.Println("\nПоиск сотрудника по имени 'Карина':")
	foundEmp, found := FindByName(employees, "Карина")
	if found {
		fmt.Println(foundEmp.Info())
	} else {
		fmt.Println("Сотрудник не найден.")
	}

	avgAge := GetAverageAge(employees)
	fmt.Printf("\nСредний возраст: %.2f лет\n", avgAge)

	fmt.Println("\nСредняя зарплата по должностям:")
	avgSalaries := GetAverageSalaries(employees)
	for position, salary := range avgSalaries {
		fmt.Printf("%s: %.2f\n", position, salary)
	}
}
