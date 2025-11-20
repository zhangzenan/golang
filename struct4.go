package main

import "fmt"

//结构体继承

// 父结构体
type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.name)
}

// 子结构体
type Dog struct {
	Age    int
	Animal //结构体嵌套 继承
	//*Animal//结构体也可以嵌套指针
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺\n", d.name)
}

func main() {
	var d = Dog{
		Age: 20,
		Animal: Animal{
			name: "阿奇",
		},
		//指针类型赋值
		// Animal: &Animal{
		// 	name: "阿奇",
		// },
	}
	d.run()
	d.wang()
}
