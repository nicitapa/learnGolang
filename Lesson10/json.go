package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string json:"name"
	Age  int    json:"age"
}

func mainfunc() {
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

	for i := range users {
		users[i].Age++
	}

	result, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при форматировании JSON:", err)
		return
	}

	fmt.Println(string(result))
}