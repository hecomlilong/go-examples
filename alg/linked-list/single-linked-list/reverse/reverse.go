package reverse

import (
	"github.com/hecomlilong/go-examples/alg/linked-list/single-linked-list/utils"
)

func Reverse(src *utils.LinkedList) *utils.LinkedList {
	var head = new(utils.LinkedList)
	var cur = src.Next
	src.Next = nil
	head.Next = src
	for cur != nil {
		var curNext = cur.Next
		var next = head.Next
		head.Next = cur
		cur.Next = next
		cur = curNext
	}
	var res = head.Next
	head = nil
	return res
}
