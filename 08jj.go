//找出整数
package main

import "fmt"

func solution(array []int) int {
	// Edit your code here
	num:=0
	i := 0
	for _,v := range array {
		if i==0{
			num=v
			i++
		}else{
			if num==v{
				i++
			}else{
				i--
			}
		}
	}
	return num
}

func main() {
	// Add your test cases here

	fmt.Println(solution([]int{1, 3, 8, 2, 3, 1, 3, 3, 3}) == 3)
}
