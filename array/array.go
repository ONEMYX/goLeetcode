package array

import (
	"fmt"
)

func Main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(test, 3)
	//s := array.SearchInsert(test, 7)
	fmt.Println(test)
}
