package main

import (
	"fmt"
	"strconv"
)

func main() {
	var y string

	fmt.Scan(&y)

	var x, _ = strconv.ParseFloat(y, 64)
	fmt.Printf("%.f", x)
}
