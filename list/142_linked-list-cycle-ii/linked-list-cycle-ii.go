package _42_linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}

	return nil
}
