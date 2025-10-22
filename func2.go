package main

import "fmt"

func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

// 递归实现1-100的和
func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}
}
func main() {

	//传入一个整数，递归打印出1-到这个数之内的所有整数
	fn1(10)

	sum := fn2(100)
	fmt.Println(sum)

}
