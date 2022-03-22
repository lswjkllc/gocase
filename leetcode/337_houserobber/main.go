package main

import "fmt"

/*
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
*/
// https://leetcode-cn.com/problems/house-robber-iii/

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

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := dp(root)

	return Max(res[0], res[1])
}

/*
dp 返回一个大小为 2 的数组 arr
arr[0] 表示不抢 root 的话, 得到的最大钱数
arr[1] 表示抢 root 的话, 得到的最大钱数
*/
func dp(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := dp(root.Left)
	right := dp(root.Right)
	// 抢 root, 下家就不能抢了
	robRoot := root.Val + left[0] + right[0]
	// 不抢 root, 下家可抢可不抢, 取决于收益大小
	notRobRoot := Max(left[0], left[1]) + Max(right[0], right[1])

	return [2]int{notRobRoot, robRoot}
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	tns1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			}},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			}}}
	fmt.Println("root:[3,2,3,null,3,null,1] => 7", rob(tns1))

	tns2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			}},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 1,
			}}}
	fmt.Println("root:[3,4,5,1,3,null,1] => 9", rob(tns2))

	tns3 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3}}}}
	fmt.Println("root:[[4,1,null,2,null,3]] => 7", rob(tns3))
}
