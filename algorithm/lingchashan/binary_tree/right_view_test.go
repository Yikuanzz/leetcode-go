package binarytree

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "空树",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "只有根节点",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: []int{1},
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
			want: []int{1, 2, 3},
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
			want: []int{1, 2, 3},
		},
		{
			name: "完整的二叉树",
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
			want: []int{1, 3, 7},
		},
		{
			name: "不规则的树",
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
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: []int{1, 3, 7},
		},
		{
			name: "多层不规则树",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 8},
						},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 6},
						Right: &TreeNode{
							Val:  7,
							Left: &TreeNode{Val: 9},
						},
					},
				},
			},
			want: []int{1, 3, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
