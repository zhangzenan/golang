package main

import (
	"fmt"
	"sync"
)

var wg4 sync.WaitGroup

func test4(num int) {
	defer wg4.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("协程（%v）打印的第%v条数据\n", num, i)
	}
}

func main() {
	for i := 0; i < 6; i++ {
		wg4.Add(1)
		go test4(i)
	}
	wg4.Wait()
	fmt.Println("关闭主线程。。。")
}
