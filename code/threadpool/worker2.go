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
	Number   int
	Tasks    chan Task
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
			default:
				time.Sleep(time.Second * 1)
			}
		}()
	}
}

func (w *Worker) Stop() {
	w.StopChan <- true
	//close(w.StopChan)
}

func main() {
	var stop chan bool
	worker := &Worker{}
	worker.RunTask(5)

	task1 := Task{F: f1, A: 1, B: 2}
	task2 := Task{F: f2, A: 3, B: 4}
	worker.AddTask(task1)
	worker.AddTask(task2)

	/*go func() {
		time.Sleep(100000 * time.Second)
		worker.Stop()	// 不会deadline
	}()*/

	for {
		select {
		//case <-worker.StopChan:
		case <-stop:
			return
		default:
			time.Sleep(time.Second * 10)
		}
	}
}
