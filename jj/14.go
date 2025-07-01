package main

import "fmt"

//小C希望构造一个包含n个元素的数组，且满足以下条件：
// 数组中的所有元素两两不同。
// 数组所有元素的最大公约数为 k。
// 数组元素之和尽可能小
func solution(n int, k int) int {
    // PLEASE DO NOT MODIFY THE FUNCTION SIGNATURE
    // write code here
	num:=0
	for i:=0;i<n;i++{
		num+=k*(i+1)
	}
    return num
	return (n * (n + 1) / 2)*k
}

func main() {
    fmt.Println(solution(3, 1) == 6)
    fmt.Println(solution(2, 2) == 6)
    fmt.Println(solution(4, 3) == 30)
}
