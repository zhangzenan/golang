package main

import "fmt"

type MyInt int

/*
*
不仅结构体能定义自定义方法，自定义类型也能定义自定义方法
*
*/
func (m MyInt) PrintInfo() {
	fmt.Println("我是自定义类型里面的方法")
}

func main() {
	var a MyInt = 20
	a.PrintInfo()
}
