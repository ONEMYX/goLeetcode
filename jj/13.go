package main

import "fmt"


//问题描述
// 小U得到了一个数字n，他的任务是构造一个特定数组。
// 这个数组的构造规则是：对于每个i从1到n，将数字n到i逆序拼接，直到i等于n为止。最终，输出这个拼接后的数组。

func solution(n int) []int {
	// write code here
	num := make([]int, 0)
	size=0
	for i:=0;i<n;i++{
		for j:=0;j<n-size;j++{
			num = append(num, n-j)
		}
	}
	return num
}


func main() {
	fmt.Println("Hello, World!")
}