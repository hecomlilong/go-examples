package main

import (
	"fmt"
)

func main() {
	readwriter()
}

func readwriter() {
	var count = 10
	go func() {
		count++
		fmt.Printf("add ...")
	}()
	go func() {
		count--
		fmt.Printf("add ...")
	}()
}
