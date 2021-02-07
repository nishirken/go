package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var arr []Person
	var fileName string

	fmt.Scan(&fileName)

	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line = scanner.Text()
		var splitted = strings.Split(line, " ")

		arr = append(arr, Person{fname: splitted[0], lname: splitted[1]})
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	file.Close()
}