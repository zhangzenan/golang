package main

import "fmt"

func main() {
	/*
		1、var 声明变量
		var 变量名称 类型

		2、一次定义多个变量
		var 变量名称,变量名称 类型

		3、一次声明多个类型不一样的变量
		var (
			username string
			age	int
			sex	string
		)

		4、定义的时候赋值
		var (
			username="张三"
			age=20
		)

		5、短变量声明法在函数内部，可以使用更简略的 := 方式声明并初始化变量
		注意：短变量只能用于声明局部变量，不能用于全局变量的声明
		短变量一次声明多个变量，并初始化
		a,b,c:=12,13,20
		短变量可以定义不同的类型
		a,b,c:=12,13,"C"

		6、匿名变量 在使用多重赋值时，如果要忽略某个值，可以使用匿名变量（anonymous variable）
		匿名变量用一个下划线_表示，例如
		var username, _ = getUserInfo()
		fmt.Println(username)

		func getUserInfo() (string, int) {
			return "zhangsan", 10
		}
		匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

		7、常量
		const username="zhangsan"
		常量必须赋值，常量的值不可以改变

		多个常量也可以一起声明
		const (
			A="A"
			B="B"
		)

		const同时声明多个常量时，如果省略了值则表示和上面的一行的值相同
		const (
			n1=100
			n2
			n3
		)

		8、iota是golang语言的常量计数器，只能在常量表达式中使用。
		iota在const关键字出现时将被重置为0（const内部的第一行之前），const每新增一行常量声明
		将使iota计数一次（iota可理解为const语句块中的行索引）。
		匿名变量_会跳过

		//iota声明中间插队
		const (
			n1 = iota
			n2 = 100
			n3 = iota
			n4
		)
		fmt.Println(n1, n2, n3, n4) //0 100 2 3

		//多个iota定义在一行
		const (
			n1, n2 = iota + 1, iota + 2
			n3, n4
			n5, n6
		)




	*/
	// var username string
	// fmt.Println(username)
	// var a1 = "张三"
	// fmt.Println(a1)
	// var a2 = 's'
	// fmt.Println(a2)

	// var a1, a2 string
	// a1 = "aaa"
	// a2 = "bbbb"
	// fmt.Println(a1, a2)

	// var (
	// 	usernaem string
	// 	age      int
	// 	sex      string
	// )
	// usernaem = "张三"
	// age = 20
	// sex = "男"
	// fmt.Println(usernaem, age, sex)

	// username := "张三"
	// fmt.Println(username)
	// fmt.Printf("类型:%T", username)

	// a, b, c := 12, 13, 20
	// fmt.Println(a, b, c)

	// var username, age = getUserInfo()
	// fmt.Println(username, age)

	// var username, _ = getUserInfo()
	// fmt.Println(username)

	// const (
	// 	n1 = 100
	// 	n2
	// 	n3
	// )
	// fmt.Println(n1, n2, n3)

	// const (
	// 	n1 = iota
	// 	n2
	// 	n3
	// 	n4
	// )
	// fmt.Println(n1, n2, n3, n4)

	//iota声明中间插队
	// const (
	// 	n1 = iota
	// 	n2 = 100
	// 	n3 = iota
	// 	n4
	// )
	// fmt.Println(n1, n2, n3, n4) //0 100 2 3

	//多个iota定义在一行
	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4
		n5, n6
	)
	fmt.Println(n1, n2)
	fmt.Println(n3, n4)
	fmt.Println(n5, n6)
}

func getUserInfo() (string, int) {
	return "zhangsan", 10
}
