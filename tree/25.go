package tree

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// 生成链表数据
func generateListNode(nums []int) *ListNode {
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	//1. 循环k个节点
	//2. 分为两个指针
	//3. 当k个节点遍历完后，将前k个节点反转
	//4. 递归下一个k个节点
	//5. 如果剩余节点不足k个，返回head
	left, right, pre := head, head.Next, head
	for i := 0; i < k-1; i++ {
		if right == nil {
			return left
		}
		right = right.Next
	}
	// 反转
	left = reverseKG(left, right)
	// 递归下一个k个节点
	pre.Next = reverseKGroup(right, k)
	return left
}

func reverseKG(head *ListNode, right *ListNode) *ListNode {
	// 链表的反转
	//第一次：
	//head = 1->2->3->4->5
	//right = 4->5
	// 第一次：
	pre := right
	cur := head
	for cur != right {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		// 第一次：
		// next = 2->3->4->5
		// cur = 1->4->5
		// pre = 1->4->5
		// cur = 2->3->4->5
		// 第二次：
		// next = 3->4->5
		// cur = 2->1->4->5
		// pre = 2->1->4->5
		// cur = 3->4->5
	}
	return pre
}
