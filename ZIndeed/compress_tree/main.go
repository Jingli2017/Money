package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	compressed := compressSparse(root)
	fmt.Println(compressed)

	r := restoreSparse(compressed)
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", r.Left)
	fmt.Printf("%#v\n", r.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compress(root *TreeNode) []int {
	// calculate tree height
	h := getTreeHeight(root)
	// create heap
	heap := make([]int, int(math.Pow(2, float64(h))))
	// BFS double queue
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)
	indexQueue := make([]int, 0)
	indexQueue = append(indexQueue, 1) // left -> 2*i right -> 2*i+1

	for len(nodeQueue) != 0 {
		currNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		currIndex := indexQueue[0]
		indexQueue = indexQueue[1:]

		heap[currIndex] = currNode.Val
		if currNode.Left != nil {
			nodeQueue = append(nodeQueue, currNode.Left)
			indexQueue = append(indexQueue, currIndex*2)
		}

		if currNode.Right != nil {
			nodeQueue = append(nodeQueue, currNode.Right)
			indexQueue = append(indexQueue, currIndex*2+1)
		}
	}

	return heap
}

func restore(heap []int) *TreeNode {
	return restoreImpl(1, heap)
}

func restoreImpl(index int, heap []int) *TreeNode {
	if index >= len(heap) {
		return nil
	}
	node := &TreeNode{Val: heap[index]}
	node.Left = restoreImpl(index*2, heap)
	node.Right = restoreImpl(index*2+1, heap)
	return node
}

// index to val
func compressSparse(root *TreeNode) map[int]int {
	indexToVal := make(map[int]int)

	indexQueue := make([]int, 0)
	indexQueue = append(indexQueue, 1)
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)

	for len(nodeQueue) != 0 {
		currNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		currIndex := indexQueue[0]
		indexQueue = indexQueue[1:]

		indexToVal[currIndex] = currNode.Val
		if currNode.Left != nil {
			nodeQueue = append(nodeQueue, currNode.Left)
			indexQueue = append(indexQueue, currIndex*2)
		}

		if currNode.Right != nil {
			nodeQueue = append(nodeQueue, currNode.Right)
			indexQueue = append(indexQueue, currIndex*2+1)
		}
	}

	return indexToVal
}

func restoreSparse(compressed map[int]int) *TreeNode {
	return restoreSparseImpl(1, compressed)
}

func restoreSparseImpl(index int, compressed map[int]int) *TreeNode {
	if _, ok := compressed[index]; !ok {
		return nil
	}

	node := &TreeNode{Val: compressed[index]}
	node.Left = restoreSparseImpl(index*2, compressed)
	node.Right = restoreSparseImpl(index*2+1, compressed)

	return node
}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMax(getTreeHeight(root.Left), getTreeHeight(root.Right)) + 1
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
