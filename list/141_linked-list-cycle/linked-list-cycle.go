package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
		if fast == head {
			return true
		}
	}
	return false

}
