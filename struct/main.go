package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(reflect.TypeOf(&a{}).Name())
	fmt.Println(reflect.TypeOf(&a{}).Elem().Name())
}

type a struct {
	a int
	b string
}
