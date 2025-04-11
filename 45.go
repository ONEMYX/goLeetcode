package main

func jump(nums []int) int {
	//nums = [2,3,1,1,4]
	//跳到最后一位的最小步数
	//1.查询步数内的最优解
	maxNum := 0
	end := 0
	cur := 0
	for i := 0; i < len(nums)-1; i++ {
		maxNum = m(maxNum, i+nums[i])
		if i == end {
			end = maxNum
			cur++
		}
	}
	return cur
}

func m(x, y int) int {
	if x > y {
		return x
	}
	return y
}
