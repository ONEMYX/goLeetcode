package array

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, op := range operations {
		if op == "++X" || op == "X++" {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
