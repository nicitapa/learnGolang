package Lesson9

import (
	"fmt"
	"time"
)

func fetch(id int) chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(id) * time.Second)
		ch <- fmt.Sprintf("Результат %d", id)
	}()
	return ch
}

func main() {
	results := []chan string{
		fetch(1),
		fetch(2),
		fetch(3),
	}

	for _, ch := range results {
		fmt.Println(<-ch)
	}
}
