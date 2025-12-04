package main

import (
	"fmt"
)

// 接口是一个规范
type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// 手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机开始工作")
}

func (c Camera) stop() {
	fmt.Println("相机停止工作")
}

func main() {
	p := Phone{
		Name: "华为手机",
	}
	var p1 Usber
	p1 = p
	p1.start()

	c := Camera{}
	var c1 Usber = c //表示相机实现了Usb接口
	c1.start()

}
