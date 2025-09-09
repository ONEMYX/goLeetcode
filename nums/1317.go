package nums

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		num := n - i
		if strings.Contains(strconv.Itoa(num), "0") || strings.Contains(strconv.Itoa(i), "0") {
			continue
		}
		return []int{i, num}
	}
	return nil
}
