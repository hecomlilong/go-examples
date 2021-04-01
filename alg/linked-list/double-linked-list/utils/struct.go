package utils

type DoubleLinkedNode struct {
	Key  int
	Val  int
	Prev *DoubleLinkedNode
	Next *DoubleLinkedNode
}

type DoubleLinkedList struct {
	head *DoubleLinkedNode
	tail *DoubleLinkedNode
	size int
}

func NewDoubleLinkedList(size int) *DoubleLinkedList {
	head, tail := new(DoubleLinkedNode), new(DoubleLinkedNode)
	head.Next = tail
	tail.Prev = head
	return &DoubleLinkedList{head: head, tail: tail, size: size}
}

func (p *DoubleLinkedList) AddLast(src *DoubleLinkedNode) {
	src.Next = p.tail
	p.tail.Prev.Next = src
	src.Prev = p.tail.Prev
	p.tail.Prev = src
	p.size++
}

func (p *DoubleLinkedList) AddFirst(src *DoubleLinkedNode) {
	src.Next = p.head.Next
	p.head.Next = src
	src.Prev = p.head
	src.Next.Prev = src
	p.size++
}

func (p *DoubleLinkedList) Remove(src *DoubleLinkedNode) {
	prev := src.Prev
	next := src.Next
	prev.Next = next
	next.Prev = prev
	p.size--
	return
}

func (p *DoubleLinkedList) RemoveFirst() *DoubleLinkedNode {
	if p.size == 0 {
		return nil
	}
	var res = p.head.Prev
	p.Remove(res)
	return res
}

func (p *DoubleLinkedList) Size() int {
	return p.size
}
