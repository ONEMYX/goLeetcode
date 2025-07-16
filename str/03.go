package str

func lengthOfLongestSubstring(s string) int {
	num := 0
	//记录每个字符的位置
	index := make(map[byte]int)
	//左指针
	left := 0
	for i := range s {
		//如果字符已经存在，更新左指针
		if lastIdx, ok := index[s[i]]; ok && lastIdx >= left {
			left = lastIdx + 1
		}
		index[s[i]] = i
		num = max(num, i-left+1)
	}
	return num
}
