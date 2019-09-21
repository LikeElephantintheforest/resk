package main

import (
	"fmt"
	"time"
)

func main() {
	testGoFunction()
}

/**
Go func 基本语法
*/
func testGoFunction() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(abc int) {
			for {
				a[abc]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
