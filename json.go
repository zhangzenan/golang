package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string //私有属性不能被json包访问
	Sno    string
}

func main() {
	//结构体转换为json字符串
	// var s1 = Student{
	// 	ID:     12,
	// 	Gender: "男",
	// 	Name:   "李四",
	// 	Sno:    "s001",
	// }
	// fmt.Printf("%#v\n", s1)

	// jsonByte, _ := json.Marshal(s1)
	// jsonStr := string(jsonByte)
	// fmt.Printf("%v", jsonStr)

	//json字符串转换为结构体对象
	var str = `{"ID":12,"Gender":"男","Name":"李四","Sno":"s001"}`
	var s1 Student
	err := json.Unmarshal([]byte(str), &s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", s1)
}
