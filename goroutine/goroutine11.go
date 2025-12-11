package main

import (
	"fmt"
	"sync"
	"time"
)

var wg11 sync.WaitGroup

// 写数据
func fn11(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) //for range循环会一直尝试从通道读取数据，只有当通道被关闭且所有数据都被读取完毕时，循环才会退出，如果通道一直不关闭，这个循环就会永远阻塞等待，报死锁错误
	wg11.Done()
}

func fn12(ch <-chan int) {
	for v := range ch {
		fmt.Printf("读取数据%v成功\n", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg11.Done()
}

func main() {
	var ch = make(chan int, 10)
	wg11.Add(1)
	go fn11(ch)

	wg11.Add(1)
	go fn12(ch)

	wg11.Wait()

	fmt.Println("执行完毕")

}
