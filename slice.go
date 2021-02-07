package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var arr []int
	var x string

	for {
		fmt.Scan(&x)

		if x == "X" {
			return
		}
		var i, _ = strconv.Atoi(x)

		arr = append(arr, i)
		sort.Ints(arr)

		fmt.Println(arr)
	}
}