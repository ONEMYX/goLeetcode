package main

import (
	"awesomeProject/array"
	"fmt"
)

func main() {

	test := []int{1, 3, 5, 6}
	s := array.SearchInsert(test, 7)
	fmt.Println(s)
}
