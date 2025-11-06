package main

import "fmt"

func main() {
	//在计算机底层 a这个变量其实对应了一个内存地址
	// var a int = 10
	// fmt.Printf("a的值：%v a的类型%T a的地址%p", a, a, &a)

	// var a = 10
	// var p = &a //p指针变量 p的类型 *int(指针类型)
	// fmt.Printf("a的值:%v a的类型%T a的地址%p\n", a, a, &a)
	// fmt.Printf("p的值:%v p的类型%T\n", p, p)

	//每一个变量都有自己的内存地址
	// var a = 10
	// var b = &a //p指针变量 p的类型 *int(指针类型)
	// fmt.Printf("a的值:%v a的类型%T a的地址%p\n", a, a, &a)
	// fmt.Printf("b的值:%v b的类型%T b的地址%p\n", b, b, &b)

	//通过指针获取值
	// a := 10
	// p := &a

	// //*p 表示取出p这个变量对应的内存地址的值
	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = 30        //改变p这个变量对应的内存地址的值
	// fmt.Println(a) //30

	// var a = 5
	// fn1(a)
	// fmt.Println(a) //5
	// fn2(&a)
	// fmt.Println(a) //40

	//通过new()创建指针
	var a = new(int) //a是一个指针变量 类型是 *int的指针类型，指针变量对应的值是0
	fmt.Printf("值:%v 类型:%T 指针变量对应的值:%v\n", a, a, *a)

	var f = new(bool)
	fmt.Println(*f)
}

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 40
}
