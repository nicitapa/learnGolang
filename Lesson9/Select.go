package Lesson9

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "данные из ch1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "данные из ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Получено:", msg)
	case msg := <-ch2:
		fmt.Println("Получено:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут: ничего не пришло вовремя")
	}
}
