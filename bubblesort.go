package main

import (
	"fmt"
)

func bublesort(arr *[]int)  {
	var i int = 0
	for i < len(*arr) {
		var j int = 0
		for j < len(*arr) - i - 1 {
			if (*arr)[j] > (*arr)[j + 1] {
				var temp int = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}

func main() {
	var arr []int
	var x string

	for i := 0; i <= 10; i++ {
		var n, _ = fmt.Scanf("%d", &x)
		arr = append(arr, n)
	}
	bublesort(&arr)
	fmt.Println(arr)
}