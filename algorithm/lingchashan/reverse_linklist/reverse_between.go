package reverselinklist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre_left := dummy
	for i := 0; i < left-1; i++ {
		pre_left = pre_left.Next
	}

	var pre, cur *ListNode = nil, pre_left.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	pre_left.Next.Next = cur
	pre_left.Next = pre
	return dummy.Next
}
