package _3__Merge_k_Sorted_Lists

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

/*
	Input: lists = [[1,4,5],[1,3,4],[2,6]]
	Output: [1,1,2,3,4,4,5,6]
	Explanation: The linked-lists are:
	[
		1->4->5,
		1->3->4,
		2->6
	]
	merging them into one sorted list:
	1->1->2->3->4->4->5->6
*/
func mergeKLists(lists []*ListNode) *ListNode {
	return sort(lists, 0, len(lists)-1)
}

func sort(lists []*ListNode, lo, hi int) *ListNode {
	if lo == hi {
		return lists[lo]
	}

	mid := (lo + hi) / 2
	l1 := sort(lists, lo, mid)
	l2 := sort(lists, mid+1, hi)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
