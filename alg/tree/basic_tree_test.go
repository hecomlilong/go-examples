package tree

import (
	"reflect"
	"testing"
)

func TestBasicTreePreOrderTraversal(t *testing.T) {
	var cases = []struct {
		Args    []interface{}
		Results []int
	}{
		{

			Args:    []interface{}{1, nil, 2, 3},
			Results: []int{1, 2, 3},
		},
		{
			Args:    []interface{}{},
			Results: []int{},
		},
		{
			Args:    []interface{}{1},
			Results: []int{1},
		},
		{
			Args:    []interface{}{1, 2},
			Results: []int{1, 2},
		},
		{
			Args:    []interface{}{1, nil, 2},
			Results: []int{1, 2},
		},
	}

	for _, item := range cases {
		t.Run("", func(t *testing.T) {
			var root = BuildATree(item.Args)
			res := PreorderTraversal(root)
			if !reflect.DeepEqual(res, item.Results) {
				t.Fatalf("got:%+v\n expect:%+v\n", res, item.Results)
			}
		})
	}
}
