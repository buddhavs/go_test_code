package main

import (
	"fmt"
	"reflect"
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
	fmt.Println(reflect.ValueOf(3).String())
}
