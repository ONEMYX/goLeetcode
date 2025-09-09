package array

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 从(1,1)开始 到角落，计算出每个点要到的数量，然后得出要计算的数量
	// 例如：(1,1) = (0,1) + (1,0)
	// (1,2) = (0,2) + (1,1)
	//（0，2） = (0,1) + (1,2)
	// (1,3) = (0,3) + (1,2)
	// (2,2) = (1,2) + (2,1)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
