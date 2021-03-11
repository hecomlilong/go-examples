package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var head = GetRandomLinkedList(10, func() int {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return r1.Intn(100)
	})
	PrintLinkedList(head, func(item *LinkedList) interface{} {
		return item.val
	})

	head = Reverse(head)

	fmt.Println("after reverse")
	PrintLinkedList(head, func(item *LinkedList) interface{} {
		return item.val
	})
}

func Reverse(src *LinkedList) *LinkedList {
	var head = new(LinkedList)
	var cur = src.next
	src.next = nil
	head.next = src
	for cur != nil {
		var curNext = cur.next
		cur.next = nil
		var next = head.next
		head.next = cur
		cur.next = next
		cur = curNext
	}
	var res = head.next
	head = nil
	return res
}

type LinkedList struct {
	val  int
	next *LinkedList
}

func NewLinkedList(val int) *LinkedList {
	var res = new(LinkedList)
	res.val = val
	return res
}

func GetRandomLinkedList(length int, valGenerator func() int) *LinkedList {
	if length <= 0 {
		return nil
	}
	var prev, head *LinkedList

	for i := 0; i < length; i++ {
		if head == nil {
			head = new(LinkedList)
			head.val = valGenerator()
			prev = head
		} else {
			prev.next = new(LinkedList)
			prev.next.val = valGenerator()
			prev = prev.next
		}
	}
	return head
}

func PrintLinkedList(head *LinkedList, itemToPrint func(item *LinkedList) interface{}) {
	var current = head
	for current != nil {
		fmt.Printf("%+v->", itemToPrint(current))
		current = current.next
	}
	fmt.Printf("NULL\n")
}
