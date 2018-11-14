package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func compute(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Printf("increment = %d\n", i)
	}
	wg.Done()
}

func main() {
	fmt.Println("before goroutine.")

	wg = &sync.WaitGroup{}

	wg.Add(1)
	go compute(5)

	wg.Add(1)
	go compute(5)

	wg.Wait()

	fmt.Println("after goroutine.")
}
