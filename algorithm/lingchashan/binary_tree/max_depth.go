package binarytree

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_1(root.Left), maxDepth_1(root.Right)) + 1
}

func maxDepth_2(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		ans = max(ans, depth)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
