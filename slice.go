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

	// var arr1 []int
	// fmt.Println(arr1)
	// //golang中声明切片以后，切片的默认值是nil
	// fmt.Println(arr1 == nil)

	//切片的循环遍历
	// var strSlice = []string{"php", "java"}
	// for i := 0; i < len(strSlice); i++ {
	// 	fmt.Println(strSlice[i])
	// }

	// for _, v := range strSlice {
	// 	fmt.Println(v)
	// }

	//基于数组定义切片
	// a := [5]int{55, 56, 57, 58, 59}
	// b := a[:] //获取数组里面的所有值
	// fmt.Printf("%v-%T\n", b, b)

	// c := a[1:4] //获取数组里面1到4
	// fmt.Printf("%v-%T\n", c, c)

	// d := a[2:] //获取数组里面索引2之后的所有
	// fmt.Printf("%v-%T\n", d, d)

	// e := a[:3] //获取第三个下标之前的数据
	// fmt.Printf("%v-%T\n", e, e)

	//基于切片定义切片
	// a := []string{"北京", "上海", "广州"}
	// b := a[1:]
	// fmt.Printf("%v-%T\n", b, b)

	//关于切片的长度和容量
	//长度：切片的长度就是它所包含的元素个数
	//容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) //长度6 容量6

	a := s[2:]
	fmt.Printf("长度%d 容量%d\n", len(a), cap(a)) //长度4 容量4

	b := s[1:3]
	fmt.Printf("长度%d 容量%d\n", len(b), cap(b)) //长度2 容量5

	c := s[:3]
	fmt.Printf("长度%d 容量%d\n", len(c), cap(c)) //长度3 容量6
	//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）

}
