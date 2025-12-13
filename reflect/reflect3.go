package main

import (
	"fmt"
	"reflect"
)

func reflectValue3(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()

	switch kind {
	case reflect.Int:
		fmt.Printf("int类型的原始值：%v\n", v.Int()+10)
	case reflect.Float32:
		fmt.Printf("float32类型的原始值：%v\n", v.Float()+10.1)
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值：%v\n", v.Float()+10)
	case reflect.String:
		fmt.Printf("String类型的原始值：%v\n", v.String()+"sss")
	default:
		fmt.Printf("还没有判断这个类型\n")

	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c string = "你好"
	reflectValue3(a)
	reflectValue3(b)
	reflectValue3(c)
}
