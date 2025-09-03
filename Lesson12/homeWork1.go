package Lesson12

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		builder.WriteString(line + " ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	text := strings.ToLower(builder.String())
	words := strings.Fields(text)

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	outFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	writer.Write([]string{"слово", "частота"})

	for word, count := range freq {
		writer.Write([]string{word, fmt.Sprintf("%d", count)})
	}

	fmt.Println("Готово! Результат сохранён в output.csv")
}
