package deletelinklist

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "空链表",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "没有重复节点的链表",
			args: args{
				head: func() *ListNode {
					// 创建链表 1->2->3->4
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					return head
				}(),
			},
			want: func() *ListNode {
				// 预期结果 1->2->3->4
				head := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node4 := &ListNode{Val: 4}
				head.Next = node2
				node2.Next = node3
				node3.Next = node4
				return head
			}(),
		},
		{
			name: "有连续重复节点的链表",
			args: args{
				head: func() *ListNode {
					// 创建链表 1->1->2->3->3
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 1}
					node3 := &ListNode{Val: 2}
					node4 := &ListNode{Val: 3}
					node5 := &ListNode{Val: 3}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node5
					return head
				}(),
			},
			want: func() *ListNode {
				// 预期结果 2
				return &ListNode{Val: 2}
			}(),
		},
		{
			name: "有多个重复节点的链表",
			args: args{
				head: func() *ListNode {
					// 创建链表 1->1->2->2->3->3->4
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 1}
					node3 := &ListNode{Val: 2}
					node4 := &ListNode{Val: 2}
					node5 := &ListNode{Val: 3}
					node6 := &ListNode{Val: 3}
					node7 := &ListNode{Val: 4}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node5
					node5.Next = node6
					node6.Next = node7
					return head
				}(),
			},
			want: func() *ListNode {
				// 预期结果 4
				return &ListNode{Val: 4}
			}(),
		},
		{
			name: "所有节点都重复的链表",
			args: args{
				head: func() *ListNode {
					// 创建链表 1->1->1->1
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 1}
					node3 := &ListNode{Val: 1}
					node4 := &ListNode{Val: 1}
					head.Next = node2
					node2.Next = node3
					node3.Next = node4
					return head
				}(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
