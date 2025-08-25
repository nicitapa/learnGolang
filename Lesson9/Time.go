package Lesson9

import (
	"fmt"
	"time"
)

func longOperation() string {
	time.Sleep(3 * time.Second)
	return "готово"
}

func main() {
	select {
	case result := <-func() chan string {
		ch := make(chan string)
		go func() {
			ch <- longOperation()
		}()
		return ch
	}():
		fmt.Println("Результат:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Прервано по таймауту")
	}

	// NewTimer
	timer := time.NewTimer(1 * time.Second)
	<-timer.C
	fmt.Println("Таймер сработал")
}
