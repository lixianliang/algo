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

var count = 0

func SortPrint() {
	//count := 1
	var wg sync.WaitGroup
	//var mutex sync.Mutex
	for i := 0; i < GoNumber; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			//var now int
			//mutex.Lock()
			//now = count
			//mutex.Unlock()
			for {
				/*mutex.Lock()
				                now = count
								                mutex.Unlock()
												                if now > Number {
																	                    return
																						                }*/
				if count > Number {
					return
				}

				//if now%GoNumber == id {
				if count%GoNumber == id {
					//fmt.Println(id, now)
					fmt.Println(id, count)
					//mutex.Lock()
					count++
					//mutex.Unlock()
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
