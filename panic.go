package main

import (
	"errors"
	"fmt"
)

//go语言中目前（go1.12）是没有异常机制，但是使用panic/recover模式来处理错误
//panic可以在任何地方引发，但recover只有在defer调用的函数中有效

func fn1() {
	fmt.Println("fn1")
}

// recover必须在defer中执行
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出一个异常")
}

// 模拟了一个读取文件的方法
func readFile(filename string) error {
	if filename == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	// fn1()
	// fn2()
	// fmt.Println("结束")

	myFn()
}
