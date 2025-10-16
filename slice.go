package main

import (
	"fmt"
	"sort"
)

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
	// s := []int{2, 3, 5, 7, 11, 13}
	// fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) //长度6 容量6

	// a := s[2:]
	// fmt.Printf("长度%d 容量%d\n", len(a), cap(a)) //长度4 容量4

	// b := s[1:3]
	// fmt.Printf("长度%d 容量%d\n", len(b), cap(b)) //长度2 容量5

	// c := s[:3]
	// fmt.Printf("长度%d 容量%d\n", len(c), cap(c)) //长度3 容量6
	// //切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）

	// var sliceA = make([]int, 4, 8)
	// fmt.Println(sliceA)
	// fmt.Printf("%d--%d\n", len(sliceA), cap(sliceA))
	// sliceA[0] = 10
	// sliceA[1] = 12
	// sliceA[2] = 40
	// sliceA[3] = 30
	// fmt.Println(sliceA)

	// sliceB := []string{"php", "java", "go"}
	// sliceB[2] = "golang"
	// fmt.Println(sliceB)

	//golang中没法通过下标的方式给切片扩容
	// var sliceC []int
	// fmt.Printf("%v -- %v  --%v", sliceC, len(sliceC), cap(sliceC))

	//golang中给切片扩容的话要用到append()方法
	// var sliceC []int
	// fmt.Printf("%v -- %v  --%v\n", sliceC, len(sliceC), cap(sliceC))
	// sliceC = append(sliceC, 12, 23, 34)
	// fmt.Printf("%v -- %v  --%v\n", sliceC, len(sliceC), cap(sliceC))

	//append方法还可以合并切片
	// sliceA := []string{"php", "java"}
	// sliceB := []string{"nodejs", "go"}

	// sliceA = append(sliceA, sliceB...)
	// fmt.Println(sliceA)

	//切片的扩容策略
	// var sliceA []int
	// for i := 1; i <= 10; i++ {
	// 	sliceA = append(sliceA, i)
	// 	fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
	// }

	/**
	值类型：改变变量副本值的时候，不会改变变量本身的值
	引用类型：改变变量副本值的时候，会改变变量本身的值
	**/
	//切片就是引用数据类型
	// sliceA := []int{1, 2, 3, 45}
	// sliceB := sliceA
	// sliceB[0] = 11
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	//copy()函数复制切片
	// sliceA := []int{1, 2, 3, 45}
	// sliceB := make([]int, 4, 4)
	// copy(sliceB, sliceA)

	// fmt.Println(sliceA)
	// fmt.Println(sliceB)
	// sliceB[0] = 111
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	//删除元素
	//Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// //要删除索引为2的元素，删除的元素：32	注意：append合并切片的时候最后一个元素要加...
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a)
	// //要删除35,36
	// sliceB := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// sliceB = append(sliceB[:5], sliceB[7:]...)
	// fmt.Println(sliceB)

	//通过切片修改字符串中的值
	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	// s2 := "你好golang"
	// runeStr := []rune(s2)
	// fmt.Println(runeStr)

	// runeStr[0] = '大'
	// fmt.Println(string(runeStr))

	//sort升序排序
	intList := []int{2, 4, 5, 7, 6, 0}
	stringList := []string{"a", "c", "z", "x"}

	sort.Ints(intList)
	sort.Strings(stringList)

	fmt.Println(intList)
	fmt.Println(stringList)

	//sort降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(intList)
	fmt.Println(stringList)

}
