package binarytree

import (
	"testing"
)

func Test_isValidBST_1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空树",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "只有根节点",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "有效的BST",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "无效的BST-左子树大于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
			},
			want: false,
		},
		{
			name: "无效的BST-右子树小于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 0},
				},
			},
			want: false,
		},
		{
			name: "复杂的有效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: true,
		},
		{
			name: "复杂的无效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 6}, // 大于根节点
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_1(tt.args.root); got != tt.want {
				t.Errorf("isValidBST_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空树",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "只有根节点",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "有效的BST",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "无效的BST-左子树大于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
			},
			want: false,
		},
		{
			name: "无效的BST-右子树小于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 0},
				},
			},
			want: false,
		},
		{
			name: "复杂的有效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: true,
		},
		{
			name: "复杂的无效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 6}, // 大于根节点
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_2(tt.args.root); got != tt.want {
				t.Errorf("isValidBST_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空树",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "只有根节点",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "有效的BST",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "无效的BST-左子树大于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
			},
			want: false,
		},
		{
			name: "无效的BST-右子树小于根节点",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 0},
				},
			},
			want: false,
		},
		{
			name: "复杂的有效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: true,
		},
		{
			name: "复杂的无效BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 6}, // 大于根节点
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_3(tt.args.root); got != tt.want {
				t.Errorf("isValidBST_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
