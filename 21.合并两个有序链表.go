/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{-101, l1}
	cur := head

	for l2 != nil && l1 != nil {
		if l2.Val <= l1.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}

	if l2 != nil {
		cur.Next = l2
	}
	if l1 != nil {
		cur.Next = l1
	}
	return head.Next
}

// @lc code=end

