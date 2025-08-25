package Lesson9

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	var atomicCounter int32

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
			atomic.AddInt32(&atomicCounter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Mutex counter:", counter)
	fmt.Println("Atomic counter:", atomicCounter)
}
