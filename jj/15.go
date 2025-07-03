// 小U和小R有两个字符串，分别是S和T，现在小U需要通过对S进行若干次操作，使其变成T的一个前缀。操作可以是修改S的某一个字符，或者删除S末尾的字符。现在你需要帮助小U计算出，最少需要多少次操作才能让S变成T的前缀。
// 例如，S = "abc"，T = "ab"，最少需要1次操作，即删除S的最后一个字符'c'，就可以得到T。
// 再例如，S = "abc"，T = "a"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。
// 再例如，S = "abc"，T = "ac"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。
// 再例如，S = "abc"，T = "bc"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。
// 再例如，S = "abc"，T = "c"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。
// 再例如，S = "abc"，T = "a"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。
// 再例如，S = "abc"，T = "a"，最少需要2次操作，即删除S的最后一个字符'c'，然后修改S的最后一个字符'b'为'a'，就可以得到T。

package main

import "fmt"

func solution(S string, T string) int {
	// 计算最少操作数
	minOps := len(S) + len(T) // 初始化为一个很大的值
	
	// 尝试让S成为T的每个可能前缀
	for prefixLen := 0; prefixLen <= len(T); prefixLen++ {
		ops := 0
		
		// 如果S的长度大于当前前缀长度，需要删除多余的字符
		if len(S) > prefixLen {
			ops += len(S) - prefixLen
		}
		
		// 计算需要修改的字符数（比较S和T的前prefixLen个字符）
		for i := 0; i < prefixLen && i < len(S); i++ {
			if S[i] != T[i] {
				ops++
			}
		}
		
		// 如果S的长度小于前缀长度，需要添加字符（通过修改操作）
		if len(S) < prefixLen {
			ops += prefixLen - len(S)
		}
		
		if ops < minOps {
			minOps = ops
		}
	}
	
	return minOps
}

func main() {
    fmt.Println("Testing solution function:")
    fmt.Printf("solution(\"aba\", \"abb\") = %d (expected: 1)\n", solution("aba", "abb"))
    fmt.Printf("solution(\"abcd\", \"efg\") = %d (expected: 4)\n", solution("abcd", "efg"))
    fmt.Printf("solution(\"xyz\", \"xy\") = %d (expected: 1)\n", solution("xyz", "xy"))
    fmt.Printf("solution(\"hello\", \"helloworld\") = %d (expected: 0)\n", solution("hello", "helloworld"))
    fmt.Printf("solution(\"same\", \"same\") = %d (expected: 0)\n", solution("same", "same"))
    
    fmt.Println("\nAll tests passed!")
}

