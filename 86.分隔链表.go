/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}

	p1, p2 := dummy1, dummy2
	p := head

	//splite
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开原始链表
		temp := p.Next
		p.Next = nil
		p = temp
	}
	//merge
	p1.Next = dummy2.Next

	return dummy1.Next

}

// @lc code=end

