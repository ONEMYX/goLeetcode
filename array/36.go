package array

// 有效数独
// 分为两部分
// 1. 整体判断每一行 每一列是否有重复数字
// 2. 每一个小九宫格是否有重复数字
func isValidSudoku(board [][]byte) bool {
	x := make(map[int]map[byte]bool)
	y := make(map[int]map[byte]bool)
	z := make(map[int]map[byte]bool)
	for i := 0; i < len(board); i++ {
		x[i] = make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			if y[j] == nil {
				y[j] = make(map[byte]bool)
			}
			if board[i][j] == '.' {
				continue
			}
			if z[i/3*3+j/3] == nil {
				z[i/3*3+j/3] = make(map[byte]bool)
			}
			if x[i][board[i][j]] || y[j][board[i][j]] || z[i/3*3+j/3][board[i][j]] {
				return false
			}

			x[i][board[i][j]] = true
			y[j][board[i][j]] = true
			z[i/3*3+j/3][board[i][j]] = true
		}
	}
	return true
}
