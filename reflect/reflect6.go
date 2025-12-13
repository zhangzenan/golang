package main

import (
	"fmt"
	"reflect"
)

type Student6 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student6) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student6) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student6) Print() {
	fmt.Println("这是一个打印方法...")
}

func PrintStructFn(s interface{}) {
	//判断参数是不是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	//1、通过类型变量里面的Method可以获取结构体的方法
	method0 := t.Method(0) //和结构体方法的顺序没有关系，和结构体方法的ACII有关系
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	//2、通过类型变量获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	//3、通过《值变量》执行方法
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	//4、执行方法传入参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) //执行方法传入参数
	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	//5、获取方法数量
	fmt.Println("方法数量：", t.NumMethod())

}

func main() {
	stu1 := Student6{
		Name:  "小明",
		Age:   15,
		Score: 30,
	}
	PrintStructFn(&stu1)

	fmt.Printf("%#v\n", stu1)
}
