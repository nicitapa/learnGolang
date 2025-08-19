package Lesson8

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go evenNumbers(&wg)

	wg.Add(1)
	go oddNumbers(&wg)
	wg.Wait()

}
func evenNumbers(wg *sync.WaitGroup) {
	fmt.Println("Четные числа ")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
	}
	wg.Done()
}
func oddNumbers(wg *sync.WaitGroup) {
	fmt.Println("Нечетные числа")
	for i := 1; i <= 9; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	wg.Done()
}
