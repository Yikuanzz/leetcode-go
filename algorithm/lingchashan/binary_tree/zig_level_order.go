package binarytree

func zigzagLevelOrder_1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	cur := []*TreeNode{root}
	level := 0
	for len(cur) > 0 {
		next := []*TreeNode{}
		vals := make([]int, len(cur))
		for i, node := range cur {
			if level%2 == 0 {
				vals[i] = node.Val
			} else {
				vals[len(cur)-i-1] = node.Val
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		ans = append(ans, vals)
		cur = next
		level++
	}
	return ans
}

func zigzagLevelOrder_2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if level%2 == 0 {
				vals[i] = node.Val
			} else {
				vals[n-i-1] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
		level++
	}
	return ans
}
