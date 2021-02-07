package main

import (
	"time"
	"fmt"
)

func sleep(d time.Duration) {
	<- time.After(d)
}

func printCurrentTime() {
	time.Sleep(time.Second * 3)
	fmt.Println("After sleep")
	fmt.Println("Current", <- time.After(time.Second * 3))
}

func main()  {
	go printCurrentTime()
	time.Sleep(time.Second * 10)
}
