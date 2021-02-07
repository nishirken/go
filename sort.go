package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
	"sort"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func f(arr *[]int, wg *sync.WaitGroup) {
	sort.Ints(*arr)
	fmt.Println(*arr)
	wg.Done()
}

func partition(arr []int) [][]int {
	var n float64 = float64(len(arr)) / 4
	fmt.Printf("Chunk size is %f\n", n)
	var temp [][]int
	for i := 0.00; i < float64(len(arr)); i += n {
		end := i + n

		if end > float64(len(arr)) {
			end = float64(len(arr))
		}

		temp = append(temp, arr[int(i):int(end)])
	}
	return temp
}

func main() {
	var len string
	fmt.Println("Enter a length of array: ")
	fmt.Scan(&len)
	var n, _ = strconv.Atoi(len)

	if n == 0 {
		fmt.Println("Empty array")
		return
	}
	
	var wg = sync.WaitGroup{}
	wg.Add(4)

	var unsorted = generateRandomSlice(n)
	fmt.Println("Initial array: ", unsorted)

	var parts = partition(unsorted)
	fmt.Println("Parts: ", parts)
	go f(&parts[0], &wg)
	go f(&parts[1], &wg)
	go f(&parts[2], &wg)
	go f(&parts[3], &wg)
	wg.Wait()
	var all []int = append(append(append(parts[0], parts[1]...), parts[2]...), parts[3]...)
	sort.Ints(all)
	fmt.Println(all)
}
