package Lesson10

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title string xml:"title"
	Year  int    xml:"year"
}

type Library struct {
	XMLName xml.Name xml:"library"
	Books   []Book   xml:"book"
}

func main() {
	data, err := os.ReadFile("books.xml")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var lib Library
	if err := xml.Unmarshal(data, &lib); err != nil {
		fmt.Println("Ошибка парсинга XML:", err)
		return
	}

	for i := range lib.Books {
		lib.Books[i].Year++
	}

	result, err := xml.MarshalIndent(lib, "", "  ")
	if err != nil {
		fmt.Println("Ошибка форматирования XML:", err)
		return
	}

	fmt.Println(string(result))
}
