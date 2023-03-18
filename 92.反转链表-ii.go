/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 var successor *ListNode

 func reverseN(head *ListNode, right int) *ListNode {
	 if right == 1 {
		 successor = head.Next
		 return head
	 }
	 last := reverseN(head.Next, right-1)
	 head.Next.Next = head
	 head.Next = successor
	 return last
 }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		// 反转1到right节点的单链表
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}
// @lc code=end

