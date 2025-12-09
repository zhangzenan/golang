package main

import "fmt"

// 循环遍历管道数据

func main() {
	//使用for range遍历通道，当通道被关闭的时候就会推出for range，如果没有关闭管道就会报错死锁错误
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭管道

	//for range循环遍历管道的值，注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}

	//通过for循环遍历管道的时候管道可以不关闭
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)

	}

}
