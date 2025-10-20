package main

import "fmt"

// 求两个数的和
func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

// 函数参数简写,如果x和y参数类型一样可以省略前面参数的类型
func subFn(x, y int) int {
	sub := x - y
	return sub
}

// 函数的可变参数，可变参数是指函数参数数量不固定，go语言中的可变参数通过在参数后加...来表示多个参数
func sumFn1(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

// return 关键词一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

// 返回值命名：函数定义时可以给返回值命名，并在函数体重直接使用这些变量，最后通过return关键字返回
func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y

	return
}

// 返回值参数类型简写
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y

	return
}

func main() {
	// sum1 := sumFn(12, 3)
	// fmt.Println(sum1)

	// sub1 := subFn(12, 3)
	// fmt.Println(sub1)

	// sum2 := sumFn1(1, 2, 3, 4)
	// fmt.Println(sum2)

	// sum, sub := calc(12, 2)
	// fmt.Println(sum)
	// fmt.Println(sub)

	// sum, sub := calc1(12, 2)
	// fmt.Println(sum)
	// fmt.Println(sub)

	// sum, sub := calc2(12, 2)
	// fmt.Println(sum)
	// fmt.Println(sub)

	sum, _ := calc2(12, 2)
	fmt.Println(sum)

}
