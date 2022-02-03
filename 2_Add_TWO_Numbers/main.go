package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p := l1
	q := l2
	var l3 *ListNode = new(ListNode)
	head := l3
	for p != nil || q != nil {
		var x int
		var y int
		if p != nil {
			x = p.Val
			p = p.Next
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
			q = q.Next
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10
		l3.Next = new(ListNode)
		l3.Next.Val = sum % 10
		l3 = l3.Next
	}
	if carry > 0 {
		l3.Next = new(ListNode)
		l3.Next.Val = carry
	}
	return head.Next
}
