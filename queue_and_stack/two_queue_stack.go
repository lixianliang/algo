package main

import (
	"container/list"
)

type Stack struct {
	queue1 *list.List
	queue2 *list.List
	// val is int
}

func NewStack() *Stack {
	return &Stack{
		queue1: list.New(),
		queue2: list.New(),
	}
}

func (s *Stack) Push(val int) {
	if s.queue1.Len() != 0 {
		s.queue1.PushBack(val)
	} else {
		s.queue2.PushBack(val)
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.queue1.Len() != 0 {
		for s.queue1.Len() > 1 {
			e := s.queue1.Front()
			s.queue2.PushBack(e.Value.(int))
			s.queue1.Remove(e)
		}
		e := s.queue1.Front()
		x := e.Value.(int)
		s.queue1.Remove(e)
		return x, true
	} else if s.queue2.Len() != 0 {
		for s.queue2.Len() > 1 {
			e := s.queue2.Front()
			s.queue1.PushBack(e.Value.(int))
			s.queue2.Remove(e)
		}
		e := s.queue2.Front()
		x := e.Value.(int)
		s.queue2.Remove(e)
		return x, true
	} else {
		return -1, false
	}
}

func main() {

}
