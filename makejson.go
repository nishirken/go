package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	var m = make(map[string]string)

	fmt.Scan(&name)
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	var encoded, _ = json.Marshal(m)

	fmt.Println(string(encoded))
}
