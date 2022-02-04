package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l3 := new(ListNode)
	head := l3
	for list1 != nil && list2 != nil {
		head.Next = new(ListNode)
		if list1.Val <= list2.Val {
			head.Next.Val = list1.Val
			list1 = list1.Next
		} else {
			head.Next.Val = list2.Val
			list2 = list2.Next
		}
		head = head.Next
	}
	for list1 != nil {
		head.Next = new(ListNode)
		head.Next.Val = list1.Val
		list1 = list1.Next
		head = head.Next
	}
	for list2 != nil {
		head.Next = new(ListNode)
		head.Next.Val = list2.Val
		list2 = list2.Next
		head = head.Next
	}
	return l3.Next

}
