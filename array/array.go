package array

func Main() {
	//test := []int{1, 2, 3, 4, 5, 6, 7}
	//test := []int{1, 2, 3, 4, 5}
	//test := []int{4, 1, 3, 3}
	//rotate(test, 3)
	//s := array.SearchInsert(test, 7)
	//num := countBadPairs(test)
	//fmt.Println(num)
	//test := []int{1, 3, 2, 3, 3}
	//countSubarrays(test, 2)
	//test := []int{2, 3, -1, 8, 4}
	//test := []int{3, 0, 8, 2, 0, 0, 1}
	//test := []int{3, 0, 6, 1, 5}
	//test := []int{1, 1, 3}
	// gas = [1,2,3,4,5], cost = [3,4,5,1,2]
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	//gas := []int{2, 3, 4}
	//cost := []int{3, 4, 3}
	canCompleteCircuit(gas, cost)
}
