package main

import (
	"fmt"
	"strconv"
)

func s() {
	//1. 建立一个通道
	one := make(chan bool)
	done := make(chan bool)
	num := make(chan bool)
	closeNum := make(chan bool)

	//add

	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Printf(strconv.Itoa(i))
				i++
				one <- true
			case <-closeNum:
				return
			}
		}
	}()

	go func() {
		for i := 'a'; i <= 'z'; i++ {
			<-one
			fmt.Printf(string(i))
			num <- true
		}
		done <- true
		closeNum <- true
	}()

	one <- true
	<-done
	fmt.Println("done")
}
