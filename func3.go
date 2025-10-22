package main

import "fmt"

/**
闭包：
1、可以让一个变量常驻内存
2、可以让一个变量不污染全局

闭包是指有权访问另一个函数作用域中的变量的函数
创建闭包的常见方式就是在一个函数内部创建另一个函数，通过另一个函数访问这个函数的局部变量
注意：由于闭包里作用域返回的局部变量资源不会被立刻销毁回收，所以可能会占用更多的内存，
过度使用闭包会导致性能下降，建议在非常有必要的情况下才使用闭包
**/

// 定义一个方法返回值也是一个方法
func adder() func() int {
	//i是局部变量，但是会常驻内存
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(int) int {
	//i是局部变量，但是会常驻内存,不污染全局
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {

	// var fn = adder()
	// fmt.Println(fn())
	// fmt.Println(fn())

	var fn2 = adder2()
	fmt.Println(fn2(1))
	fmt.Println(fn2(1))

}
