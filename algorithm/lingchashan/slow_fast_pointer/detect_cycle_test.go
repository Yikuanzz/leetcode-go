package slowfastpointer

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "空链表",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "单个节点",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: nil,
		},
		{
			name: "无环链表",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "有环链表",
			args: args{
				head: func() *ListNode {
					// 创建环形链表 1->2->3->4->2
					node1 := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					node1.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node2
					return node1
				}(),
			},
			want: func() *ListNode {
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node4 := &ListNode{Val: 4}
				node2.Next = node3
				node3.Next = node4
				node4.Next = node2
				return node2
			}(),
		},
		{
			name: "两个节点的环形链表",
			args: args{
				head: func() *ListNode {
					// 创建环形链表 1->2->1
					node1 := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node1.Next = node2
					node2.Next = node1
					return node1
				}(),
			},
			want: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node1.Next = node2
				node2.Next = node1
				return node1
			}(),
		},
		{
			name: "环在链表末尾",
			args: args{
				head: func() *ListNode {
					// 创建环形链表 1->2->3->4->5->3
					node1 := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					node5 := &ListNode{Val: 5}
					node1.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node5
					node5.Next = node3
					return node1
				}(),
			},
			want: func() *ListNode {
				node3 := &ListNode{Val: 3}
				node4 := &ListNode{Val: 4}
				node5 := &ListNode{Val: 5}
				node3.Next = node4
				node4.Next = node5
				node5.Next = node3
				return node3
			}(),
		},
		{
			name: "自环链表",
			args: args{
				head: func() *ListNode {
					// 创建自环链表 1->1
					node1 := &ListNode{Val: 1}
					node1.Next = node1
					return node1
				}(),
			},
			want: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node1.Next = node1
				return node1
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
