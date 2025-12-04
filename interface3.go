package main

import "fmt"

// x.(type)判断一个变量的类型，这个语句只能用在switch语句里面
func Myprint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	default:
		fmt.Println("传入错误")
	}
}

// 类型断言
func main() {
	var a interface{}
	a = "你好"
	v, ok := a.(string)
	if ok {
		fmt.Printf("a是一个string 类型，值：%v\n", v)
	} else {
		fmt.Println("断言失败")
	}

	v1, ok := a.(bool)
	if ok {
		fmt.Printf("a是一个string 类型，值：%v\n", v1)
	} else {
		fmt.Println("断言失败")
	}

	Myprint("你好")
	Myprint(10)
}
