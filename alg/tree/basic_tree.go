package tree

import "github.com/hecomlilong/go-examples/alg/utils"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	var res []int
	var stack = utils.NewStack(100)
	var cur = root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
		pop, err := stack.Pop()
		if err != nil {
			continue
		}
		cur = pop.(*TreeNode).Right
	}
	return res
}

func InorderTraversal(root *TreeNode) []int {
	var res []int
	var stack = utils.NewStack(100)
	var cur = root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		pop, err := stack.Pop()
		if err != nil {
			continue
		}
		item, ok := pop.(*TreeNode)
		if !ok {
			continue
		}
		res = append(res, item.Val)
		cur = item.Right
	}
	return res
}
