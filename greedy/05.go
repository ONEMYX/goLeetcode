package main

func longestPalindrome(s string) string {
	// 最长回文字段 中心扩展法
	// 1. 取中心位置 注意 中心位置 可能是 一个字符 也可能是 两个字符
	// 2. 向两边扩展
	// 3. 比较左右两边的字符是否相等
	// 4. 如果相等 则继续扩展
	// 5. 如果不相等 则返回当前的回文子串的长度
	// 6. 比较当前回文子串的长度 是否大于 最大回文子串的长度
	// 7. 如果大于 则更新最大回文子串的长度
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 1. 取中心位置 注意 中心位置 可能是 一个字符 也可能是 两个字符
		l1, r1 := comparePalindrome(s, i, i)
		l2, r2 := comparePalindrome(s, i, i+1)
		// 6. 先比较单个字段和双字段的回文子串长度
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]

}

// 字段比较
// 2. 向两边扩展
// 3. 比较左右两边的字符是否相等
// 4. 如果相等 则继续扩展
// 5. 如果不相等 则返回当前的回文子串的长度

func comparePalindrome(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
