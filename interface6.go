package main

import "fmt"

//一个结构体实现多个接口

type Animaler1 interface {
	SetName(string)
}
type Animaler2 interface {
	GetName() string
}
type Dog1 struct {
	Name string
}

func (d *Dog1) SetName(name string) {
	d.Name = name
}

func (d Dog1) GetName() string {
	return d.Name
}

func main() {
	//Dog实现Animal的接口
	d := &Dog1{
		Name: "小黑",
	}
	var d1 Animaler1 = d
	var d2 Animaler2 = d

	d1.SetName("阿奇111")
	fmt.Println(d2.GetName())

}
