package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
		select {
		case <-time.After(1 * time.Nanosecond):
			return
		default:
		}
	}
}
