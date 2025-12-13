package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	// v.Name() //类型名称
	// v.Kind() //种类
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e)
	reflectFn(f)

	var h = 34
	reflectFn(&h)

	var i = [3]int{1, 2, 3}
	reflectFn(i)

	var j = []int{11, 2, 33}
	reflectFn(j)

}
