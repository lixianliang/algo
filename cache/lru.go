package main

import (
	"container/list"
	"fmt"
)

type Lru struct {
	cap  int
	list *list.List
	kvs  map[int]*list.Element
}

type Node struct {
	key, val int
}

func NewLru(k int) *Lru {
	return &Lru{
		cap:  k,
		list: list.New(),
		kvs:  make(map[int]*list.Element, k),
	}
}

func (this *Lru) Get(key int) int {
	e, ok := this.kvs[key]
	if ok {
		// 存在
		this.list.MoveToFront(e)
		return e.Value.(*Node).val
	}

	return -1
}

func (this *Lru) Set(key, val int) {
	e, ok := this.kvs[key]
	if ok {
		// 已存在
		if val != e.Value.(*Node).val {
			//	e.Value = val
			e.Value.(*Node).val = val
		}
		this.list.MoveToFront(e)
	} else {
		// 不存在
		if this.list.Len() >= this.cap {
			// 超出容量
			e = this.list.Back()
			delete(this.kvs, e.Value.(*Node).key) // 要删除key
			this.list.Remove(e)
		}

		e = this.list.PushFront(&Node{key: key, val: val})
		this.kvs[key] = e
	}
}

func LRU(operators [][]int, k int) []int {
	// write code here
	ret := []int{}
	lru := NewLru(k)
	for _, arr := range operators {
		if arr[0] == 1 {
			lru.Set(arr[1], arr[2])
		} else if arr[0] == 2 {
			ret = append(ret, lru.Get(arr[1]))
		}
	}
	return ret
}

func main() {
	ops := [][]int{{1, 1, 1}, {1, 1, 2}, {2, 1}}
	ret := LRU(ops, 3)
	fmt.Println(ret)
}
