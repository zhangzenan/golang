package main

import "fmt"

func main() {
	//1、创建channel
	//chan管道的关键字，int管道类型，3表示管道容量
	ch := make(chan int, 3)

	//2、给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3、获取管道里面的内容
	a := <-ch
	fmt.Println(a)

	<-ch //从管道里面获取值
	c := <-ch
	fmt.Println(c)
	ch <- 56
	ch <- 66

	//4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度：%v\n", ch, cap(ch), len(ch)) //值：0xc0000b8000 容量：3 长度：2

	//5、管道的类型（引用数据类型）
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d)

	//8、管道阻塞
	ch6 := make(chan int, 1)
	ch6 <- 34
	//ch6 <- 64 //如果管道的容量放不下数据就会报错：atal error: all goroutines are asleep - deadlock!

	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报错
	ch7 := make(chan string, 2)
	ch7 <- "数据1"
	ch7 <- "数据2"
	m1 := <-ch7
	m2 := <-ch7
	m3 := <-ch7
	fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

}
