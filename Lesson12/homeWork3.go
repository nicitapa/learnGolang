package Lesson12

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	f := excelize.NewFile()
	sheet := "Sheet1"

	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Age")

	for i, user := range users {
		row := i + 2 // начинаем со второй строки
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Age)
	}

	if err := f.SaveAs("report.xlsx"); err != nil {
		fmt.Println("Ошибка при сохранении Excel:", err)
		return
	}

	fmt.Println("Файл report.xlsx успешно создан!")
}
