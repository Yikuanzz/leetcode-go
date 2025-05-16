package binarytree

import (
	"testing"
)

func Test_maxDepth_1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "空树",
			args: args{root: nil},
			want: 0,
		},
		{
			name: "只有根节点",
			args: args{root: &TreeNode{Val: 1}},
			want: 1,
		},
		{
			name: "平衡二叉树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: 3,
		},
		{
			name: "不平衡二叉树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			}},
			want: 4,
		},
		{
			name: "只有左子树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth_1(tt.args.root); got != tt.want {
				t.Errorf("maxDepth_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth_2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "空树",
			args: args{root: nil},
			want: 0,
		},
		{
			name: "只有根节点",
			args: args{root: &TreeNode{Val: 1}},
			want: 1,
		},
		{
			name: "平衡二叉树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: 3,
		},
		{
			name: "不平衡二叉树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			}},
			want: 4,
		},
		{
			name: "只有左子树",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth_2(tt.args.root); got != tt.want {
				t.Errorf("maxDepth_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
