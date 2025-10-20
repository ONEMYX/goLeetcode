package array

func findSmallestInteger(nums []int, value int) int {
	//1. 记录每个数对value的余数
	//2. 记录每个余数的出现次数
	//3. 从0开始遍历，找到第一个出现次数为0的余数
	remainderCount := make(map[int]int)
	for _, num := range nums {
		n := ((num % value) + value) % value
		remainderCount[n]++
	}
	i := 0
	for remainderCount[i%value] > 0 {
		remainderCount[i%value]--
		i++
	}
	return i
}
