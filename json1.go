package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签
type Student1 struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	//结构体转换为json字符串
	// var s1 = Student1{
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
	var str = `{"id":12,"gender":"男","name":"李四","sno":"s001"}`
	var s1 Student1
	err := json.Unmarshal([]byte(str), &s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", s1)
}
