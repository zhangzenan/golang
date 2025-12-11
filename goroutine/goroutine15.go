package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
var wg15 sync.WaitGroup
var mutex1 sync.RWMutex

// 写的方法
func write() {
	mutex1.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex1.Unlock()
	wg15.Done()
}

func read() {
	mutex1.RLock()
	fmt.Println("----执行读操作")
	time.Sleep(time.Second * 2)
	mutex1.RUnlock()
	wg15.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg15.Add(1)
		go write()
		wg15.Add(1)
		go read()
	}
	wg15.Wait()
}
