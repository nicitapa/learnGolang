package Lesson9

import (
	"fmt"
	"time"
)

func unbuffered() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Отправка:", i)
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Получено:", val)
	}
}

func buffered() {
	ch := make(chan int, 2)

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Буфер отправка:", i)
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Буфер получено:", val)
	}
}

func main() {
	fmt.Println("=== Небуферизированный канал ===")
	unbuffered()
	time.Sleep(time.Second)

	fmt.Println("\n=== Буферизированный канал ===")
	buffered()
}
