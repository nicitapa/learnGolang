package Lesson12

import (
	"bufio"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Программа запущена. Введите строки (exit — для выхода):")

	for {
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if text == "exit" {
			log.Println("Программа завершена по команде exit.")
			break
		}
		log.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка при чтении ввода: %v", err)
	}
}
