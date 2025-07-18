package matrix

// 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	//1. 记录方向
	//2. 记录是否撞墙
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1
	result := make([]int, 0, rows*cols)
	for left <= right && top <= bottom {
		// 从左到右
		for cols = left; cols <= right; cols++ {
			result = append(result, matrix[left][cols])
		}
		top++
		// 从上到下
		for rows = top; rows <= bottom; rows++ {
			result = append(result, matrix[rows][right])
		}
		right--
		if top <= bottom {
			// 从右到左
			for cols = right; cols >= left; cols-- {
				result = append(result, matrix[bottom][cols])
			}
			bottom--
		}

		if left <= right {
			// 从下到上
			for rows = bottom; rows >= top; rows-- {
				result = append(result, matrix[rows][left])
			}
			left++
		}

	}
	return result
}
