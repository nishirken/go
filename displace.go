package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a float64, v float64, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (a * t * t) / 2 + v * t + s
	}
}

func main() {
	var a string
	var v string
	var s string
	var t string

	fmt.Println("Acceleration: ")
	fmt.Scan(&a)

	fmt.Println("Velocity: ")
	fmt.Scan(&v)

	fmt.Println("Displacement: ")
	fmt.Scan(&s)

	fmt.Println("Time: ")
	fmt.Scan(&t)

	var a1, _ = strconv.ParseFloat(a, 64)
	var v1, _ = strconv.ParseFloat(v, 64)
	var s1, _ = strconv.ParseFloat(s, 64)
	var t1, _ = strconv.ParseFloat(t, 64)

	f := GenDisplaceFn(a1, v1, s1)

	fmt.Println(f(t1))
}