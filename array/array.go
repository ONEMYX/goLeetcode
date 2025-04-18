package array

import "fmt"

func Main() {
	//test := []int{1, 2, 3, 4, 5, 6, 7}
	test := []int{1, 2, 3, 4, 5}
	//test := []int{4, 1, 3, 3}
	//rotate(test, 3)
	//s := array.SearchInsert(test, 7)
	num := countBadPairs(test)
	fmt.Println(num)
}
