package array

// gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// gas 加的油  cost 到达需要的油
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	for i := 0; i < length; i++ {
		// 油箱的油
		if gas[i] < cost[i] {
			continue
		}
		fuel := 0
		for k := 0; k <= length; k++ {
			if k == length {
				return i
			}

			key := (i + k) % (length)
			fuel = fuel + gas[key] - cost[key]
			if fuel < 0 {
				break
			}

		}

	}

	return -1
}
