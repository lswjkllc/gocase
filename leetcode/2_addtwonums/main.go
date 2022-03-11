package main

import "fmt"

/*
题⽬⼤意
	2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。
解题思路
	需要注意的是各种进位问题。
	为了处理⽅法统⼀，可以先建⽴⼀个虚拟头结点，这个虚拟头结点的 Next 指向真正的 head。
		这样 head 不需要单独处理，直接 while 循环即可。
		另外判断循环终⽌的条件不⽤是 p.Next ！= nil，这样最后⼀位还需要额外计算，循环终⽌条件应该是 p != nil。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := &ListNode{Val: 0, Next: nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		carry = (x + y + carry) / 10
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: &ListNode{Val: 5, Next: nil}}}
	fmt.Println(AddTwoNumbers(l1, l2))

}
