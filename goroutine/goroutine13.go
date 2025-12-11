package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello word")
	}
}
func test13() {
	//这里我们可以使用defer+recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error,没有分配空间

}

func main() {
	go sayHello()
	go test13()

	time.Sleep(time.Millisecond * 500)
}
