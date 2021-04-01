package utils

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func NewLinkedList(val int) *LinkedList {
	var res = new(LinkedList)
	res.Val = val
	return res
}

func GetRandomLinkedList(length int, ValGenerator func() int) *LinkedList {
	if length <= 0 {
		return nil
	}
	var prev, head *LinkedList

	for i := 0; i < length; i++ {
		if head == nil {
			head = new(LinkedList)
			head.Val = ValGenerator()
			prev = head
		} else {
			prev.Next = new(LinkedList)
			prev.Next.Val = ValGenerator()
			prev = prev.Next
		}
	}
	return head
}

func PrintLinkedList(head *LinkedList, itemToPrint func(item *LinkedList) interface{}) {
	var current = head
	for current != nil {
		fmt.Printf("%+v->", itemToPrint(current))
		current = current.Next
	}
	fmt.Printf("NULL\n")
}
func ConvertLinkedListToSlice(src *LinkedList) []int {
	var cur = src
	var res []int
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func CheckReverse(src, dst []int) bool {
	if len(src) != len(dst) {
		return false
	}
	l := len(src)
	for i, v := range src {
		if v != dst[l-i-1] {
			return false
		}
	}
	return true
}
