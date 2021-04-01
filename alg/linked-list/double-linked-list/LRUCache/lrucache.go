package lrucache

import (
	"errors"

	"github.com/hecomlilong/go-examples/alg/linked-list/double-linked-list/utils"
)

type LRUCache struct {
	cacheMap   map[int]*utils.DoubleLinkedNode
	doubleList *utils.DoubleLinkedList
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		cacheMap:   make(map[int]*utils.DoubleLinkedNode),
		doubleList: utils.NewDoubleLinkedList(size),
	}
}

func (p *LRUCache) Put(key, val int) {
	p.AddRecently(key, val)
}

func (p *LRUCache) Get(key int) (int, error) {
	if v, ok := p.cacheMap[key]; ok {
		p.AddRecently(key, v.Val)
		return v.Val, nil
	} else {
		return 0, errors.New("not found")
	}
}

func (p *LRUCache) AddRecently(key, val int) {

}

func (p *LRUCache) Remove(key int) {

}
