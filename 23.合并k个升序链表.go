/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(a *ListNode, b *ListNode, head *ListNode) *ListNode {
	if a == nil || b == nil {
		if a == nil {
			return b
		} else {
			return a
		}
	}
	head.Next = a
	cur := head
	for a != nil && b != nil {
		if b.Val <= a.Val {
			cur.Next = b
			cur = cur.Next
			b = b.Next
		} else {
			cur.Next = a
			cur = cur.Next
			a = a.Next
		}
	}
	if b != nil {
		cur.Next = b
	}
	if a != nil {
		cur.Next = a
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{-10000, nil}
	head := &ListNode{-10000, nil}
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i], head)
	}
	return ans.Next
}

// @lc code=end

