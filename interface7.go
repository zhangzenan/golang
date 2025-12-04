package main

import "fmt"

//接口嵌套接口

type Ainterface interface {
	SetName(string)
}
type Binterface interface {
	GetName() string
}

// 接口的嵌套
type Animaler3 interface {
	Ainterface
	Binterface
}
type Dog2 struct {
	Name string
}

func (d *Dog2) SetName(name string) {
	d.Name = name
}

func (d Dog2) GetName() string {
	return d.Name
}

func main() {
	//Dog实现Animal的接口
	d := &Dog2{
		Name: "小黑",
	}
	var d1 Animaler3 = d

	d1.SetName("阿奇111")
	fmt.Println(d1.GetName())

}
