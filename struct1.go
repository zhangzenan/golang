package main

import "fmt"

/**
在go语言中，没有类的概念但是可以给类型（结构体，自定义类型）定义方法。所谓方法就是
定义了接收者的函数。接收者的概念就类似于其他语言中的this或者self
方法的定义格式如下：
func(接收者变量 接收者类型) 方法名（参数列表）（返回参数）{
	函数体
}
golang中结构体是相互独立的，不会相互影响
**/

type Person1 struct {
	Name   string
	Age    int
	Sex    string
	Height int
}

func (p Person1) PringInfo() {
	fmt.Printf("姓名：%v 年龄：%v\n", p.Name, p.Age)
}

func (p *Person1) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	var p1 = Person1{
		Name: "张三",
		Age:  20,
		Sex:  "男",
	}
	p1.PringInfo()

	var p2 = Person1{
		Name: "李四",
		Age:  20,
		Sex:  "男",
	}
	p2.PringInfo()
	p1.PringInfo()
	p1.SetInfo("王五", 34)
	p1.PringInfo()
}
