package deletelinklist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "删除倒数第一个节点",
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
				n: 1,
			},
			want: func() *ListNode {
				// 预期结果 1->2->3
				head := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				head.Next = node2
				node2.Next = node3
				return head
			}(),
		},
		{
			name: "删除倒数第二个节点",
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
				n: 2,
			},
			want: func() *ListNode {
				// 预期结果 1->2->4
				head := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node4 := &ListNode{Val: 4}
				head.Next = node2
				node2.Next = node4
				return head
			}(),
		},
		{
			name: "删除头节点",
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
				n: 4,
			},
			want: func() *ListNode {
				// 预期结果 2->3->4
				head := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node4 := &ListNode{Val: 4}
				head.Next = node3
				node3.Next = node4
				return head
			}(),
		},
		{
			name: "只有一个节点的链表",
			args: args{
				head: &ListNode{Val: 1},
				n:    1,
			},
			want: nil,
		},
		{
			name: "空链表",
			args: args{
				head: nil,
				n:    1,
			},
			want: nil,
		},
		{
			name: "删除倒数第三个节点",
			args: args{
				head: func() *ListNode {
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
					return head
				}(),
				n: 3,
			},
			want: func() *ListNode {
				// 预期结果 1->2->4->5
				head := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node4 := &ListNode{Val: 4}
				node5 := &ListNode{Val: 5}
				head.Next = node2
				node2.Next = node4
				node4.Next = node5
				return head
			}(),
		},
		{
			name: "删除倒数最后一个节点",
			args: args{
				head: func() *ListNode {
					// 创建链表 1->2->3
					head := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					head.Next = node2
					node2.Next = node3
					return head
				}(),
				n: 3,
			},
			want: func() *ListNode {
				// 预期结果 2->3
				head := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				head.Next = node3
				return head
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
