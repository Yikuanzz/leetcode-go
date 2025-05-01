package slowfastpointer

func reverseLinklist(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := middleNode(head)
	head2 := reverseLinklist(mid)
	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}
