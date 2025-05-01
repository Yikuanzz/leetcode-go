package slowfastpointer

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
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
			name: "单个节点",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "两个节点",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "奇数个节点",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
		{
			name: "偶数个节点",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !reflect.DeepEqual(tt.args.head, tt.want) {
				t.Errorf("reorderList() = %v, want %v", tt.args.head, tt.want)
			}
		})
	}
}
