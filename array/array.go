package main

import (
	"errors"
	"fmt"
)

type Array struct {
	last int
	cap  int
	arr  []int
}

func NewArray(size int) *Array {
	return &Array{
		last: 0,
		cap:  size,
		arr:  make([]int, size),
	}
}

func (a *Array) Push(val int) {
	if a.last == a.cap {
		a.cap *= 2
		newarr := make([]int, a.cap)
		copy(newarr, a.arr)
		a.arr = newarr
	}

	a.arr[a.last] = val
	a.last++
}

func (a *Array) Get(index int) (int, error) {
	if index >= a.last {
		return 0, errors.New("out of index")
	}

	return a.arr[index], nil
}

func main() {
	arr := NewArray(2)
	arr.Push(3)
	arr.Push(4)
	arr.Push(5)

	fmt.Println(arr)
}
