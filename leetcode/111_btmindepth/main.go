package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthWithBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 1

	for queue != nil {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}

	return depth
}

func minDepthWithDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepthWithDfs(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepthWithDfs(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	// [1,2,3,4,null,null,5]
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	bfsD1, dfsD1 := minDepthWithBfs(root1), minDepthWithDfs(root1)
	fmt.Println(bfsD1 == dfsD1, bfsD1, dfsD1)
	// [3,9,20,null,null,15,7]
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	bfsD2, dfsD2 := minDepthWithBfs(root2), minDepthWithDfs(root2)
	fmt.Println(bfsD2 == dfsD2, bfsD2, dfsD2)
}
