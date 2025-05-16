package binarytree

import "testing"

func Test_isSymmetric(t *testing.T) {
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
			name: "对称的树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: true,
		},
		{
			name: "不对称的树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
				},
			},
			want: false,
		},
		{
			name: "结构对称但值不对称的树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5}, // 值不同
					},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
