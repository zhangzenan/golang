package main

import "fmt"

func main() {
	//切片的声明 初始化
	// var arr1 []int
	// fmt.Printf("%v - %T -长度:%v\n", arr1, arr1, len(arr1))

	// var arr2 = []int{1, 2, 34, 45}
	// fmt.Printf("%v - %T -长度:%v\n", arr2, arr2, len(arr2))

	// var arr3 = []int{1: 2, 2: 4, 5: 34, 8: 45}
	// fmt.Printf("%v - %T -长度:%v\n", arr3, arr3, len(arr3))

	//关于nil的认识
	//当你声明了一个变量，但却还没有赋值时，golang中会自动给你的变量赋值一个默认零值。
	// 这是每种类型对应的零值：
	/**
	bool->false
	numbers->0
	string->""
	pointers->nil
	slices->nil
	maps->nil
	channels->nil
	functions->nil
	**/

	var arr1 []int
	fmt.Println(arr1)
	//golang中声明切片以后，切片的默认值是nil
	fmt.Println(arr1 == nil)

}
