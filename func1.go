package main

import "fmt"

// 函数作为另一个函数的参数
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type calcType func(int, int) int

func calc(x, y int, cb calcType) int {
	return cb(x, y)
}

//返回值是一个函数

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	// sum := calc(10, 5, add)
	// fmt.Println(sum)

	// //匿名函数作为参数
	// j := calc(3, 4, func(x, y int) int {
	// 	return x * y
	// })
	// fmt.Println(j)

	//方法作为返回值
	// var a = do("+")
	// fmt.Printf("a的类型:%T\n", a)
	// fmt.Println(a(12, 4))

	//匿名函数 匿名函数自执行函数
	func() {
		fmt.Println("test ...")
	}()

	//匿名函数
	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 3))

	//匿名自执行函数接收参数
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}
