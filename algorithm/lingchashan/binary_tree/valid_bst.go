package binarytree

import "math"

func isValidBST_1(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		x := node.Val
		return x > min && x < max && dfs(node.Left, min, x) && dfs(node.Right, x, max)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func isValidBST_2(root *TreeNode) bool {
	var inorder func(node *TreeNode) bool
	var pre int = math.MinInt
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}

func isValidBST_3(root *TreeNode) bool {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt
		}
		lMin, lMax := dfs(node.Left)
		rMin, rMax := dfs(node.Right)
		x := node.Val
		if lMax >= x || rMin <= x {
			return math.MinInt, math.MaxInt
		}
		return min(lMin, x), max(rMax, x)
	}
	_, mx := dfs(root)
	return mx != math.MinInt
}
