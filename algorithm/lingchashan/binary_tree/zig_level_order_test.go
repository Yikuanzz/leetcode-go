package binarytree

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder_1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "空树测试",
			args: args{root: nil},
			want: [][]int{},
		},
		{
			name: "单节点树测试",
			args: args{root: &TreeNode{Val: 1}},
			want: [][]int{{1}},
		},
		{
			name: "普通二叉树测试",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "复杂二叉树测试",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			}},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder_1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zigzagLevelOrder_2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "空树测试",
			args: args{root: nil},
			want: [][]int{},
		},
		{
			name: "单节点树测试",
			args: args{root: &TreeNode{Val: 1}},
			want: [][]int{{1}},
		},
		{
			name: "普通二叉树测试",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "复杂二叉树测试",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			}},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder_2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
