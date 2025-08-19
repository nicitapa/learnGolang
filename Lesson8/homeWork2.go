package Lesson8

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	t := time.Now()
	fmt.Println("getting started")
	wg.Add(1)
	go loadFirstFile(&wg)

	wg.Add(1)
	go loadSecondFile(&wg)

	wg.Add(1)
	go loadThirdFile(&wg)

	wg.Wait()
	fmt.Println("Program exited.")
	fmt.Println(time.Since(t))

}

func loadFirstFile(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("file 1 loaded")
	wg.Done()
}
func loadSecondFile(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("file 2 loaded")
	wg.Done()
}
func loadThirdFile(wg *sync.WaitGroup) {
	time.Sleep(1 / 2 * time.Second)
	fmt.Println("file 3 loaded")
	wg.Done()
}
