package tree

import "fmt"

func BuildATree(src []interface{}) *TreeNode {
	var root *TreeNode
	if len(src) == 0 {
		return root
	}
	var m = make(map[int]*TreeNode)

	for i, v := range src {
		if v != nil {
			var cur = &TreeNode{Val: v.(int)}
			m[i] = cur
			fmt.Println("i:", i, "v:", v, "parent:", (i-1)/2)
			if i == 0 {
				continue
			} else {
				var parent = (i - 1) / 2
				if i%2 != 0 {
					m[parent].Left = cur
				} else {
					m[parent].Right = cur
				}
			}
		}
	}

	return m[0]
}
