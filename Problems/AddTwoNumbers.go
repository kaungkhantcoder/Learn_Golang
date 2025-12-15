package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}

	current := dummyHead

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}

		sum := val1 + val2 + carry

		carry = sum / 10

		digit := sum % 10

		current.Next = &ListNode{Val: digit}

		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummyHead.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func AddTwoNumbers() {
	l1a := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2a := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	resultA := addTwoNumbers(l1a, l2a)
    fmt.Print("Input 1: ")
    printList(l1a) // 2 -> 4 -> 3
    fmt.Print("Input 2: ")
    printList(l2a) // 5 -> 6 -> 4
    fmt.Print("Result:  ")
    printList(resultA) // 7 -> 0 -> 8 (Represents 807)
    
    fmt.Println("---")

    // Example 2: (9) + (9 -> 9)
    // Represents: 9 + 99 = 108
    l1b := &ListNode{Val: 9}
    l2b := &ListNode{Val: 9, Next: &ListNode{Val: 9}}

    resultB := addTwoNumbers(l1b, l2b)
    fmt.Print("Input 1: ")
    printList(l1b) // 9
    fmt.Print("Input 2: ")
    printList(l2b) // 9 -> 9
    fmt.Print("Result:  ")
    printList(resultB) // 8 -> 0 -> 1 (Represents 108)
}