package main

import (
	"sync"
)

var i = 0
var once = &sync.Once{}

func GetISomewhere() int {
	return 1
}

func CreateInstance() {
  once.Do(func() {
    i = GetISomewhere()
  })
}
