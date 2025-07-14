package str

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	num := 0
	for i < j {
		h1, h2 := height[i], height[j]
		num = max(num, min(h1, h2)*(j-i))
		if h1 > h2 {
			j--
		} else {
			i++
		}
	}
	return num
}
