package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg14 sync.WaitGroup
var mutex sync.Mutex

func test14() {
	mutex.Lock()
	count++
	fmt.Println("the count is:", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg14.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg14.Add(1)
		go test14()
	}
	wg14.Wait()
}
