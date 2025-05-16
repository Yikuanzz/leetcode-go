package binarytree

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "空树",
			args: args{
				root: nil,
				p:    &TreeNode{Val: 1},
				q:    &TreeNode{Val: 2},
			},
			want: 0,
		},
		{
			name: "根节点是LCA",
			args: func() args {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				return args{
					root: root,
					p:    root.Left,
					q:    root.Right,
				}
			}(),
			want: 3,
		},
		{
			name: "左子树中的LCA",
			args: func() args {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 6}
				root.Left.Right = &TreeNode{Val: 2}
				return args{
					root: root,
					p:    root.Left.Left,
					q:    root.Left.Right,
				}
			}(),
			want: 5,
		},
		{
			name: "右子树中的LCA",
			args: func() args {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Right.Left = &TreeNode{Val: 0}
				root.Right.Right = &TreeNode{Val: 8}
				return args{
					root: root,
					p:    root.Right.Left,
					q:    root.Right.Right,
				}
			}(),
			want: 1,
		},
		{
			name: "一个节点是另一个节点的祖先",
			args: func() args {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 6}
				root.Left.Right = &TreeNode{Val: 2}
				return args{
					root: root,
					p:    root.Left,
					q:    root.Left.Left,
				}
			}(),
			want: 5,
		},
		{
			name: "复杂的树结构",
			args: func() args {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 5}
				root.Right = &TreeNode{Val: 1}
				root.Left.Left = &TreeNode{Val: 6}
				root.Left.Right = &TreeNode{Val: 2}
				root.Left.Right.Left = &TreeNode{Val: 7}
				root.Left.Right.Right = &TreeNode{Val: 4}
				root.Right.Left = &TreeNode{Val: 0}
				root.Right.Right = &TreeNode{Val: 8}
				return args{
					root: root,
					p:    root.Left.Right.Left,
					q:    root.Left.Right.Right,
				}
			}(),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q)
			if tt.args.root == nil {
				if got != nil {
					t.Errorf("lowestCommonAncestor() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Errorf("lowestCommonAncestor() = nil, want %v", tt.want)
				return
			}
			if got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
