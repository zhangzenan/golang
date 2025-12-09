package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：使用goroutine和channel协同工作案例
1、开启一个WriteData的协程给向管道inChan中写入100条数据
2、开启一个ReadData的协程读取inChan中写入的数据
3、注意：WriteData和ReadData同时操作一个管道
4、主线程必须等待操作完成后才可以退出
goroutine结合Channel使用的简单demo,定义两个方法，一个方法给管道里面写数据，一个给管道里面读数据
*/

var wg8 sync.WaitGroup

// 写数据
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) //for range循环会一直尝试从通道读取数据，只有当通道被关闭且所有数据都被读取完毕时，循环才会退出，如果通道一直不关闭，这个循环就会永远阻塞等待，报死锁错误
	wg8.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("读取数据%v成功\n", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg8.Done()
}

func main() {
	var ch = make(chan int, 10)
	wg8.Add(1)
	go fn1(ch)

	wg8.Add(1)
	go fn2(ch)

	wg8.Wait()

	fmt.Println("执行完毕")

}
