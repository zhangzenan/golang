package main

import (
	"fmt"
	"sync"
)

var wg9 sync.WaitGroup

// 向intChan放入1-120000个数
func putNum(intChan chan int) {
	for i := 2; i < 100; i++ {
		intChan <- i
	}
	close(intChan)
	wg9.Done()
}

// 从intChan取出数据，并判断是否为素数，如果是，就把得到的素数放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num //num是素数
		}
	}
	//primeChan
	//close(primeChan) //如果一个channel关闭了就没法给这个channel发送数据了
	//什么时候关闭primeChan？

	//给exitChan里面放入一条数据
	exitChan <- true

	wg9.Done()

}

// printPrime打印素数的方法
func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg9.Done()
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16) //标识primeChan close

	//存放数字的协程
	wg9.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0; i < 16; i++ {
		wg9.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//打印素数的协程
	wg9.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满
	wg9.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg9.Done()
	}()
	wg9.Wait()

	fmt.Println("执行完毕")
}
