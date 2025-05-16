package binarytree

func levelOrder_1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		nxt := []*TreeNode{}
		vals := make([]int, len(cur))
		for i, node := range cur {
			vals[i] = node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		ans = append(ans, vals)
		cur = nxt
	}
	return ans
}

func levelOrder_2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n)
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
