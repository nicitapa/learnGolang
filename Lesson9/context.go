package Lesson9

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Работа завершена")
	case <-ctx.Done():
		fmt.Println("Прервано:", ctx.Err())
	}
}

func main() {
	// Timeout
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()
	doWork(ctx1)

	// Cancel
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel2()
	}()
	doWork(ctx2)
}
