package Lesson9

import (
	"fmt"
	"sync"
)

var count int

func race() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			count++ // гонка!
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Без защиты:", count)
}

func fixedWithMutex() {
	count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("С Mutex:", count)
}

func main() {
	race()
	fixedWithMutex()
}
