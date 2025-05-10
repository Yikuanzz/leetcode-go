package deletelinklist

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "删除中间节点",
			args: args{
				node: func() *ListNode {
					// 创建链表 1->2->3->4
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					return node2 // 要删除的节点
				}(),
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		{
			name: "删除倒数第二个节点",
			args: args{
				node: func() *ListNode {
					// 创建链表 1->2->3->4
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					return node3 // 要删除的节点
				}(),
			},
			want: &ListNode{
				Val: 4,
			},
		},
		{
			name: "删除倒数第三个节点",
			args: args{
				node: func() *ListNode {
					// 创建链表 1->2->3->4->5
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					node5 := &ListNode{Val: 5}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node5
					return node2 // 要删除的节点
				}(),
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if !reflect.DeepEqual(tt.args.node, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}
}
