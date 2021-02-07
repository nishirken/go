package main

import (
	"fmt"
	"time"
)

func decrement(x *int) {
	*x--
}

func increment(x *int) {
	*x++
}

// x value depends on what function inc or decr will call first
func main() {
	var x int = 0
	go decrement(&x)
	go increment(&x)
	time.Sleep(1000)
	fmt.Println(x)
}
