package main

import "fmt"

// 单向管道
func main() {

	//1、在默认情况下，管道是双向的
	ch10 := make(chan int, 2)
	ch10 <- 10
	ch10 <- 12
	m10 := <-ch10
	m11 := <-ch10
	fmt.Println(m10, m11)

	//2、管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	//<-ch2

	//3、管道声明为只读
	ch3 := make(<-chan int, 2)
	//ch3 <- 23

}
