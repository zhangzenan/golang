package main

import (
	"fmt"
	"sync"
	"time"
)

/*
在主线程（可以理解成进程）中，开启一个goroutine,该协程每隔1秒输出你好golang
在主线程中也每隔1秒输出你好golang,输出10次后，退出程序
要求主线程和goroutine同时执行
*/

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1():你好golang-", i)
		time.Sleep(time.Millisecond * 1000)
	}
	wg.Done() //协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2():你好golang-", i)
		time.Sleep(time.Millisecond * 1000)
	}
	wg.Done() //协程计数器-1
}

func main() {
	wg.Add(1)  //协程计数器加1
	go test1() //表示开启一个协程

	wg.Add(1)
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Println("main()你好golang-", i)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Wait()
	fmt.Println("主线程退出")
}
