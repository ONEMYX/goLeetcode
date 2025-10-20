package main

import (
	"awesomeProject/array"
)

func main() {
	array.Main()
	//backtracking.Main()
	//str.Main()
	//list := generateList([]int{1, 2})
	//removeNthFromEnd(list, 2)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 生成链表
func generateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here

	// 1. 两个指针 一个先走 N 然后同步走
	// 2. 当 r 指针。next== nil 时  l 进行删除操作

	left := head
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	if right == nil {
		return head.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	next := left.Next
	left.Next = next.Next
	return head
}
