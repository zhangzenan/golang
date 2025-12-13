package main

import (
	"fmt"
	"reflect"
)

type Student7 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student7) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student7) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student7) Print() {
	fmt.Println("这是一个打印方法...")
}

func reflectChangeStruce(s interface{}) {
	//判断参数是不是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体指针类型")
		return
	}
	//修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("李四")
	age := v.Elem().FieldByName("Age")
	age.SetInt(22)

}

func main() {
	stu1 := Student7{
		Name:  "小明",
		Age:   15,
		Score: 30,
	}
	reflectChangeStruce(&stu1)

	fmt.Printf("%#v\n", stu1)
}
