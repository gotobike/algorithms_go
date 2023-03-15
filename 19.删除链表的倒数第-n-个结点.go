/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	slow1, slow2, fast := dummy, head, head

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow1 = slow1.Next
		slow2 = slow2.Next
	}

	slow1.Next = slow2.Next

	// if slow2 != head {

	// } else {
	// 	slow1.Next = nil
	// }

	return dummy.Next

}

// @lc code=end

