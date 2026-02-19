package main

import (
	"fmt"
	"time"
	"sync"
)

func processOrder(id int) {
	fmt.Printf("Processing order with id: %d\n", id)
	time.Sleep(2 * time.Second)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processOrder(i + 1)
		}()
	}
	wg.Wait()
}