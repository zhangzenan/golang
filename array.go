package main

import "fmt"

func main() {
	//数组的长度是类型的一部分
	// var arr1 [3]int
	// var arr2 [4]int
	// var strArr [3]string
	// fmt.Printf("arr1:%T arr2:%T strArr:%T", arr1, arr2, strArr)

	//数组的初始化，第一种方法
	// var arr1 [3]int
	// fmt.Println(arr1)
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 3
	// fmt.Println(arr1)

	//数组的初始化，第二种方法
	// var arr1 = [3]int{23, 34, 5}
	// fmt.Println(arr1)

	// arr1 := [3]int{23, 34, 5}
	// fmt.Println(arr1)

	//数组的初始化，第三种方法 一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
	// var arr1 = [...]int{1, 2, 3, 4}
	// fmt.Println(arr1)

	//数组的初始化，第四种方法，指定下标初始化
	// arr := [...]int{0: 1, 1: 10, 2: 20, 5: 50}
	// fmt.Println(len(arr))
	// fmt.Println(arr)

	//数组的循环遍历，for	for range
	// var arr = [3]int{23, 34, 5}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// var arr = [3]int{23, 34, 5}
	// for k, v := range arr {
	// 	fmt.Printf("key:%v  value:%v\n", k, v)
	// }

	//值类型	引用类型
	//基本数据类型	和	数组都是值类型
	// var a = 10
	// b := a
	// a = 20
	// fmt.Println(a, b)

	// var arr1 = [...]int{1, 2, 3}
	// arr2 := arr1
	// arr1[0] = 11
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	//引用类型：切片
	// var arr1 = []int{1, 2, 3}
	// arr2 := arr1
	// arr1[0] = 11
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	//多维数组
	// var arr = [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(arr[0][1])
	// for _, v1 := range arr {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	//定义数组的另一种方式，多维数组只有第一层可以 使用...来让编译器推导数组长度
	var arr = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr[0][1])

}
