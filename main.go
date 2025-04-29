package main

import "fmt"

func main() {
	//array.Main()
	list := new(TreeNode)
	list.Val = 4
	list.Left = new(TreeNode)
	list.Left.Val = 9
	list.Left.Left = new(TreeNode)
	list.Left.Left.Val = 5
	list.Left.Right = new(TreeNode)
	list.Left.Right.Val = 1
	list.Right = new(TreeNode)
	list.Right.Val = 0
	s := sumNumbers(list)
	fmt.Println(s)
}
