package binarytree

import "math"

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}

	// 检查左右子树高度差是否超过1
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}
