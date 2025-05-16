package binarytree

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	x := root.Val
	if p.Val < x && q.Val < x {
		if root.Left == nil {
			return root
		}
		return lowestCommonAncestorBST(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		if root.Right == nil {
			return root
		}
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	return root
}
