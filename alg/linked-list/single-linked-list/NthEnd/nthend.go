package nthend

import (
	"github.com/hecomlilong/go-examples/alg/linked-list/single-linked-list/utils"
)

func NthEnd(src *utils.LinkedList, n int) *utils.LinkedList {
	var left, right *utils.LinkedList = src, src

	var i int
	for i = 1; i < n && right.Next != nil; i++ {
		right = right.Next
	}
	if i < n {
		return nil
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
