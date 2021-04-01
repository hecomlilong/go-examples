package reverse

import (
	"math/rand"
	"testing"
	"time"

	"github.com/hecomlilong/go-examples/alg/linked-list/single-linked-list/utils"
)

func TestReverse(t *testing.T) {
	var head = utils.GetRandomLinkedList(10, func() int {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return r1.Intn(100)
	})

	src := utils.ConvertLinkedListToSlice(head)

	head = Reverse(head)
	dst := utils.ConvertLinkedListToSlice(head)
	if !utils.CheckReverse(src, dst) {
		t.Fatalf("src:%+v\ndst:%+v\n", src, dst)
	}
}
