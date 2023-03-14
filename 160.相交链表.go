/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}

	}

	return pa
	// lenA, lenB := 0, 0

	// for pa := headA; pa.Next != nil; pa = pa.Next {
	// 	lenA++
	// }
	// for pb := headB; pb.Next != nil; pb = pb.Next {
	// 	lenB++
	// }

	// pa, pb := headA, headB

	// if lenA > lenB {
	// 	for i := 0; i < lenA-lenB; i++ {
	// 		pa = pa.Next
	// 	}
	// } else {
	// 	for i := 0; i < lenB-lenA; i++ {
	// 		pb = pb.Next
	// 	}
	// }

	// for pa != pb {
	// 	pa = pa.Next
	// 	pb = pb.Next
	// }

	// return pa
}

// @lc code=end

