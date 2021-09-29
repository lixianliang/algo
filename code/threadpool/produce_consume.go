package main

import (
	"fmt"
	"time"
)

type Worker struct {
	Msg chan string
}

func (w *Worker) Init(size int) {
	w.Msg = make(chan string, size)
}

func (w *Worker) Consume(n int) {
	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case msg := <-w.Msg:
					fmt.Println("Consume recv: ", msg)
				case <-time.After(time.Second * 3):
				}
			}
		}()
	}
}

func (w *Worker) Produce(n int) {
	for i := 0; i < n; i++ {
		go func(number int) {
			for j := 0; j < 5; j++ {
				w.Msg <- fmt.Sprintf("number g:%d count:%d", number, j)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func main() {
	worker := &Worker{}
	worker.Init(10)
	worker.Consume(5)
	worker.Produce(5)

	for {
		time.Sleep(time.Second * 3)
	}
}
