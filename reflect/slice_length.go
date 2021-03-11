package main

import (
	"log"
	"reflect"
)

func SomeKindOfSlice() interface{} {
	return []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}
func SomeKindOfArray() interface{} {
	return [10]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}
func main() {
	log.Println(reflect.ValueOf(SomeKindOfSlice()).Len())
	log.Println(reflect.ValueOf(SomeKindOfSlice()).Kind())
	log.Println(reflect.ValueOf(SomeKindOfArray()).Type().Len())
}
