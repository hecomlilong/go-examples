package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(squareForArrPipeline(1, 100))
}
func squareForArrPipeline(min, max int) []int {
	var res []int = make([]int, max-min+1)
	in := echo(makeRange(min, max))

	var fs [3]<-chan int
	for i := range fs {
		fs[i] = square(in)
	}
	var i = 0
	for c := range merge(fs[:]) {
		res[i] = c
		i++
	}
	return res
}

func squareForArray(min, max int) []int {
	var res []int = make([]int, max-min+1)
	for i := 0; i < max-min+1; i++ {
		res[i] = (min + i) * (min + i)
	}
	return res
}

func echo(in []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range in {
			out <- n
		}

		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func makeRange(min, max int) []int {
	var res = make([]int, max-min+1)
	for i := 0; i < max-min+1; i++ {
		res[i] = i + min
	}
	return res
}

func merge(ins []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	for i := range ins {
		go func(i int) {
			for c := range ins[i] {
				out <- c
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
