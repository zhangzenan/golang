package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

func PrintStructField(s interface{}) {
	//判断参数是不是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	//1、通过类型变量里面的Field可以获取结构体的字段
	field0 := t.Field(0)
	fmt.Printf("%#v \n", field0)
	fmt.Println("字段名称:", field0.Name)
	fmt.Println("字段类型:", field0.Type)
	fmt.Println("字段Tag:", field0.Tag.Get("json"))

	//2、通过类型变量里面的FieldByName可以获取结构体的字段
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称:", field1.Name)
		fmt.Println("字段类型:", field1.Type)
		fmt.Println("字段Tag:", field1.Tag.Get("json"))
	}

	//3、通过类型变量里面的NunField获取到该结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//4、通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 30,
	}
	PrintStructField(stu1)
}
