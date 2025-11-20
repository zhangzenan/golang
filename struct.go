package main

import "fmt"

// 名称首字母是大写表示公有的，首字母小写表示私有的
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	//第一种实例化方法
	var p1 Person //实例化Person结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 20
	fmt.Printf("值：%v 类型:%T\n", p1, p1)
	//%#v打印详细信息
	fmt.Printf("值：%#v 类型:%T\n", p1, p1)

	//第二种实例化方法
	//注意：在Golang中支持对结构体指针直接使用.来访问结构体成员，p2.name="张三"其实在底层是(*p2).name
	var p2 = new(Person)
	p2.name = "李四"
	p2.age = 20
	p2.sex = "男"
	(*p2).name = "王五"
	fmt.Printf("值：%v 类型:%T\n", p2, p2)
	//%#v打印详细信息
	fmt.Printf("值：%#v 类型:%T\n", p2, p2)

	//第三种方式,通过地址方式实例化
	var p3 = &Person{}
	p3.name = "李四"
	p3.age = 25
	p3.sex = "男"
	fmt.Printf("值：%#v 类型:%T\n", p3, p3)

	//第四种方式
	var p4 = Person{
		name: "哈哈",
		age:  15,
		sex:  "男",
	}
	fmt.Printf("值：%#v 类型:%T\n", p4, p4)

	//第五种方式
	var p5 = &Person{
		name: "哈哈",
		age:  15,
		sex:  "男",
	}
	fmt.Printf("值：%#v 类型:%T\n", p5, p5)

	//第六种方式
	var p6 = &Person{
		name: "哈哈",
	}
	fmt.Printf("值：%#v 类型:%T\n", p6, p6)

	//第七种方式
	var p7 = &Person{
		"张三",
		20,
		"男",
	}
	fmt.Printf("值：%#v 类型:%T\n", p7, p7)

}
