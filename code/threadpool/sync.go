package main

import (
	"fmt"
	"sync"
	//  "time"
)

const (
	GoNumber = 8
	Number   = 85
)

func SortPrint() {
	count := 1
	var wg sync.WaitGroup
	//var mutex sync.Mutex
	for i := 0; i < GoNumber; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				if count > Number {
					return
				}

				if count%GoNumber == id {
					fmt.Println(id, count)
					count++
				}
				//time.Sleep(time.Millisecond * 1)
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	SortPrint()
}
