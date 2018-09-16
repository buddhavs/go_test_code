package main

import (
	"fmt"

	_ "github.com/buddhavs/go_test_code"
)

type Run struct{}

func (r *Run) Jump() {}

type I interface {
	Jump()
}

func main() {
	/*
		[:] is not deep copy.
	*/
	slice := []int{1, 2, 3}
	copySlice := slice[:]
	slice[0] = 42
	fmt.Println(copySlice)

	/*

	 */
	slice = make([]int, 3, 3)
}
