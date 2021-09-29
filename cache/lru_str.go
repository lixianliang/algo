package main

// 必须定义一个 包名为 `main` 的包，并实现 `main()` 函数。

import (
	"container/list"
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
}

type LRUCache interface {
	Put(key string, value string)
	Get(key string) (value string, err error)
	Delete(key string)
}

func New(maxNum int) LRUCache {
	return &Lru{
		Len:   maxNum,
		Kvs:   make(map[string]*list.Element, maxNum),
		Lists: list.New(),
	}
}

type Node struct {
	Key string
	Val string
}

type Lru struct {
	Len   int
	Kvs   map[string]*list.Element
	Lists *list.List
}

func (lru *Lru) Put(key, value string) {
	// 查找是否存在
	e, ok := lru.Kvs[key]
	if ok {
		node := e.Value.(*Node)
		/*if node.Val == value {
			return
		}*/

		if node.Val != value {
			node.Val = value
		}
		e.Value = node
		lru.Lists.MoveToFront(e)
	} else {
		if lru.Lists.Len() >= lru.Len {
			e = lru.Lists.Back()
			lru.Lists.Remove(e)
			delete(lru.Kvs, e.Value.(*Node).Key)
			//delete(lru.Kvs, key)
		}

		node := &Node{Key: key, Val: value}
		e = lru.Lists.PushFront(node)
		lru.Kvs[key] = e
	}
}

func (lru *Lru) Get(key string) (value string, err error) {
	e, ok := lru.Kvs[key]
	if ok {
		node := e.Value.(*Node)
		lru.Lists.MoveToFront(e)
		return node.Val, nil
	} else {
		return "", errors.New("not found")
	}
}

func (lru *Lru) Delete(key string) {
	e, ok := lru.Kvs[key]
	if ok {
		delete(lru.Kvs, key)
		lru.Lists.Remove(e)
	}
}
