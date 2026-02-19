package main

import (
	"fmt"
	"time"
	"sync"
)

var a = 0
var mtx = sync.Mutex{}

func Add() {
	mtx.Lock()
	defer mtx.Unlock()
	a++
}

func main() {
	for i := 0; i < 500; i++ {
		go Add()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(a)
}