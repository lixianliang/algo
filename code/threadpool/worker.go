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
	Tasks  chan Task
	Number int
}

func (w *Worker) Init(number int) {
	w.Tasks = make(chan Task, number)
	w.Number = number
}

func (w *Worker) AddTask(t Task) {
	w.Tasks <- t
}

func (w *Worker) RunTask() {
	//var wg sync.WaitGroup
	for i := 0; i < w.Number; i++ {
		go func() {
			select {
			case task := <-w.Tasks:
				//go task.F(task.A, task.B)
				fmt.Println(task.F(task.A, task.B))
			default:
				time.Sleep(time.Second * 1)
			}
		}()
	}
}

var stop chan bool

func main() {
	worker := &Worker{}
	worker.Init(5)
	go worker.RunTask()

	task1 := Task{F: f1, A: 1, B: 2}
	task2 := Task{F: f2, A: 3, B: 4}
	worker.AddTask(task1)
	worker.AddTask(task2)

	//<-stop
	for {
		time.Sleep(time.Second * 1)
	}
}
