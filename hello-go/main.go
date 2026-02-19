package main

import (
	"fmt"
	"sync"
)

var results = make(chan float64, 10)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(h float64, w float64) {
			defer wg.Done() 
			r := Rectangle {
				height: h,
				width: w,
			}
			results <- r.Area()
		}(float64(i + 2), float64(i + 1))
	}

	// Kỹ thuật quan trọng: Chạy một Goroutine riêng để đợi và đóng channel
	// Nếu không đóng channel, vòng lặp range ở dưới sẽ bị deadlock
	go func() {
		wg.Wait()
		close(results) 
	}()
	
	var totalArea float64

	for area := range results {
		totalArea += area
	}

	fmt.Printf("Total area: %.2f\n", totalArea)
}