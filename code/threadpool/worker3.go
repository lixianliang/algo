package main

import (
	"fmt"
	"time"
)

func f1(a, b int) int {
	return a + b
}

func f2(a, b int) int {
	return a * b
}

type Proc func(int, int) int

type Task struct {
	F    Proc
	A, B int
}

type Worker struct {
	StopChan chan bool
	//	Number   int
	Tasks chan Task
}

func (w *Worker) AddTask(t Task) {
	w.Tasks <- t
}

func (w *Worker) RunTask(number int) {
	w.StopChan = make(chan bool)
	w.Tasks = make(chan Task, number)
	for i := 0; i < number; i++ {
		go func() {
			select {
			case task := <-w.Tasks:
				fmt.Println(task.F(task.A, task.B))
			case <-time.After(time.Second * 3):
				fmt.Println("not get task")
				// 提供可推出机会
				//default:
				//	time.Sleep(time.Second * 1)
			}
		}()
	}
}

func main() {
	worker := &Worker{}
	worker.RunTask(8)

	task1 := Task{F: f1, A: 1, B: 2}
	task2 := Task{F: f2, A: 3, B: 4}
	worker.AddTask(task1)
	worker.AddTask(task2)

	// worker不需要提供loop循环，让其服务自己提供

	for {
		time.Sleep(time.Second * 3)
	}
}
