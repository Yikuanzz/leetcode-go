package binarytree

import "testing"

func Test_isBalanced(t *testing.T) {
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
			name: "平衡的树",
			args: args{
				root: &TreeNode{
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
				},
			},
			want: true,
		},
		{
			name: "不平衡的树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:  3,
							Left: &TreeNode{Val: 4},
						},
					},
					Right: &TreeNode{Val: 5},
				},
			},
			want: false,
		},
		{
			name: "只有左子树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: false,
		},
		{
			name: "只有右子树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: false,
		},
		{
			name: "复杂的不平衡树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:  4,
								Left: &TreeNode{Val: 5},
							},
						},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:  7,
						Left: &TreeNode{Val: 8},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
