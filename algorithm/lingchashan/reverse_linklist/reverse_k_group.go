package reverselinklist

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	var p, pre, cur *ListNode = dummy, nil, head
	for length >= k {
		length -= k
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		tmp := p.Next
		p.Next.Next = cur
		p.Next = pre
		p = tmp
	}
	return dummy.Next
}
