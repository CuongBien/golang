// Khi write, các read và write khác sẽ bị blocked.
// Khi read, write sẽ bị blocked
// Các read không blocked lẫn nhau
package main

import (
	"fmt"
	"sync"
)

var a = 0
var mtx = sync.RWMutex{}

func Add() {
	mtx.Lock()
	defer mtx.Unlock()
	a++
}

func Get() int {
	mtx.RLock()
	defer mtx.RUnlock()
	return a
}

func main() {
	for i := 0; i < 500; i++ {
		go Get()
	}
	fmt.Println(a)
}