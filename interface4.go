package main

import "fmt"

/*
结构体指针接收者实现接口：
指针接收者：如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量,
结构体值类型没法赋值给接口变量
*/

type Usber1 interface {
	start()
	stop()
}

type Phone1 struct {
	Name string
}

// 指针接收者
func (p *Phone1) start() {
	fmt.Println(p.Name, "启动")
}

func (p *Phone1) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	var phone1 = &Phone1{
		Name: "小米",
	}
	var p1 Usber1 = phone1
	p1.start()
}
