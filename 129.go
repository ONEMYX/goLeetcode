package main

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

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, prevSum int) int {
	if node == nil {
		return 0
	}
	sum := prevSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return dfs(node.Left, sum) + dfs(node.Right, sum)
}
