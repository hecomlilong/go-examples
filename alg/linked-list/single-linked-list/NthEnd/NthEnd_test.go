package nthend

import (
	"math/rand"
	"testing"
	"time"

	"github.com/hecomlilong/go-examples/alg/linked-list/single-linked-list/utils"
)

func TestNthEnd(t *testing.T) {

	var head = utils.GetRandomLinkedList(10, func() int {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return r1.Intn(100)
	})

	src := utils.ConvertLinkedListToSlice(head)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := 1 + r1.Intn(10)

	var node = NthEnd(head, n)
	if node == nil {
		t.Fatalf("src:%+v\nnode:%+v\n", src, node)
	}
	t.Logf("src:%+v   node.Val:%+v   n:%+v", src, node.Val, n)

	if src[len(src)-n] != node.Val {
		t.Errorf("node.Val:%+v\n", node.Val)
		t.Fatalf("src:%+v\nnode:%+v\n", src, node)
	}
}
