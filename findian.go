package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scan(&x)

	var lower string = strings.ToLower(x)

	var isIAN = strings.Contains(lower, "a") && strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n")

	if isIAN {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
