package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	//var num = 10 + x//mismatched types int and interface{}

	//通过断言的方式
	b, _ := x.(int)
	var num = 10 + b
	fmt.Println(num)

	//反射来实现这个功能
	v := reflect.ValueOf(x)
	fmt.Println(v)

	var n = v.Int() + 12
	fmt.Println(n)

}

func main() {
	reflectValue(10)

}
