package main

import "fmt"

func f1() {
	fmt.Println("开始")

	defer func() {
		defer fmt.Println("aaaaa")
		defer fmt.Println("bbbbb")
	}()

	fmt.Println("结束")
}

// 调用方法返回的是0
func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func main() {

	//被defer定义的语句会延迟执行，最先定义defer的语句会最后执行，下面执行结果 开始 结束 3 2 1
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")

	//f1()

	//fmt.Println(f2())
	fmt.Println(f3())
}
