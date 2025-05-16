package binarytree

import (
	"reflect"
	"testing"
)

func Test_levelOrder_1(t *testing.T) {
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
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder_1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder_2(t *testing.T) {
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
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder_2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
