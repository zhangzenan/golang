package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())

	fmt.Println(v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("你好go")
	}
}

func main() {
	var a int64 = 100
	reflectSetValue(&a)

	fmt.Println(a)
}
