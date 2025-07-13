package str

func romanToInt(s string) int {
	n := len(s)
	num := 0
	Values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k := range s {
		value := Values[s[k]]
		//边界问题
		if k < n-1 && value < Values[s[k+1]] {
			num -= value
		} else {
			num += value
		}
	}
	return num
}
