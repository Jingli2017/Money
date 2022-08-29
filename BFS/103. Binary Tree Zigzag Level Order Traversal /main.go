package _03__Binary_Tree_Zigzag_Level_Order_Traversal_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	isFromLeft := true

	for len(q) != 0 {
		temp := make([]int, 0)
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if isFromLeft {
				temp = append(temp, node.Val)
			} else {
				temp = insert(temp, 0, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		isFromLeft = !isFromLeft
		res = append(res, temp)
	}

	return res
}

func insert(nums []int, index, val int) []int {
	if index == len(nums) {
		nums = append(nums, val)
		return nums
	}

	nums = append(nums[:index+1], nums[index:]...)
	nums[index] = val
	return nums
}
