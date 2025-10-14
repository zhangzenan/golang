package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、定义int类型
	// var num int = 10
	// num = 22
	// fmt.Printf("num=%v 类型:%T", num, num)

	//2、int8的范围演示(-128到127)
	// var num int8 = 98
	// fmt.Printf("num=%v 类型:%T", num, num)

	//unsafe.Sizeof 查看不同长度的整型，在内存里面的存储空间

	// var a int8 = 15
	// fmt.Printf("num=%v 类型:%T\n", a, a)
	// fmt.Println(unsafe.Sizeof(a)) //单位字节

	//3、uint8的范围（0-255）
	// var n1 uint8 = 10
	// fmt.Printf("num=%v 类型:%T\n", n1, n1)

	// int类型转换
	// var a1 int32 = 10
	// var a2 int64 = 21
	// fmt.Println(a1 + int32(a2))
	// fmt.Println(int64(a1) + a2)

	//数字字面量语法 %v原样输出 %d表示10进制输出 %b表示二进制输出 %o表示八进制输出 %x表示十六进制输出
	// num := 30
	// fmt.Println("num=%v 类型:%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //表示64位的计算机 int就是int64 占用8个字节

	//2、定义float类型
	// var a float32 = 3.12
	// fmt.Printf("值:%v--%f,类型%T\n", a, a, a) //float32占4个字节
	// fmt.Println(unsafe.Sizeof(a))

	//%f输出float类型	%.2f 输出数据的时候保留2位小数
	// var c float64 = 3.1415926
	// fmt.Printf("%v--%f--%.2f\n", c, c, c)

	//3、64位系统中，浮点数默认是float64
	// f1 := 3.134556677
	// fmt.Printf("%f--%T", f1, f1)

	//4、Golang科学计数法表示浮点类型
	// var f2 = 3.14e2 //表示f2等于3.14*10的二次方
	// fmt.Println(f2)

	// var f3 float32 = 3.14e-2 //0.0314表示3.14除以10的二次方
	// fmt.Printf("%v--%T", f3, f3)

	//bool
	//布尔类型变量的默认值为false
	//go语言中不允许将整型强制转换为布尔类型
	//布尔类型无法参与数值运算，也无法与其他类型进行转换
	// var flag = true
	// fmt.Printf("%v---%T", flag, flag)

	//多行字符串
	// str1 := `this is str
	// this is str
	// this is str
	// `
	// fmt.Println(str1)

	//字符串操作
	//len(str)求长度
	// var str1 = "你好"
	// fmt.Println(len(str1))

	//+或者fmt.Springf拼接字符串
	// str1 := "你好"
	// str2 := "golang"
	// fmt.Println(str1 + str2)

	// str1 := "你好"
	// str2 := "golang"
	// str3 := fmt.Sprintf("%v %v", str1, str2)
	// fmt.Println(str3)

	//strings.Split 分割字符串	strings需要引入strings包
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// fmt.Println(arr) //[123 456 789]切片（简单的理解切片就是数组 在golang中切片和数组还有区别）

	//strings.Join()	join操作 表示把切片链接成字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")

	// str2 := strings.Join(arr, "*")
	// fmt.Println(str2)

	//strings.contains判断是否包含
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.Contains(str1, str2)
	// fmt.Println(flag)

	//strings.HasPrefix   前缀判断
	//strings.HasSuffix		后缀判断
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1, str2)
	// fmt.Println(flag)
	// str1 := "this is str"
	// str2 := "str"
	// flag := strings.HasSuffix(str1, str2)
	// fmt.Println(flag)

	//strings.Index()	strings.LastIndex()子串出现的位置
	// str1 := "this is str"
	// str2 := "str"
	// num := strings.LastIndex(str1, str2)
	// fmt.Println(num)

	//golang中定义字符，字符属于int类型
	// var a = 'a'
	// fmt.Printf("值：%v 类型：%T", a, a)

	//原样输出字符
	// var a = 'a'
	// fmt.Printf("值：%c 类型：%T", a, a)

	//定义一个字符串输出字符串里面的字符
	// var str = "this"
	// fmt.Printf("值：%v 原样输出%c 类型：%T", str[1], str[1], str[1])

	//一个汉字占用3个字节（utf-8）,一个字母占用一个字节

	//通过循环输出字符串里面的字符
	// s := "你好 golang"
	// for i := 0; i < len(s); i++ { //byte
	// 	fmt.Printf("%v(%c)", s[i], s[i])
	// }

	// s := "你好 golang"
	// for _, v := range s {
	// 	fmt.Printf("%v(%c)\n", v, v)
	// }

	//修改字符串
	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	// s1 := "你好golang"
	// runeStr := []rune(s1)
	// runeStr[0] = '大'
	// fmt.Println(string(runeStr))

	//整型和整型之间的转换
	// var a int8 = 20
	// var b int16 = 40
	// fmt.Println(int16(a) + b)

	//浮点型和浮点型之间的转换
	// var a float32 = 20
	// var b float64 = 40
	// fmt.Println(float64(a) + b)

	//通过fmt.Sprintf()把其他类型转换成String类型
	//注意：Sprintf使用中需要注意转换的格式 int为%d	float为%f	bool为%t	byte为%c
	// var i int = 20
	// var f float64 = 12.456
	// var t bool = true
	// var b byte = 'a'
	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("值：%v 类型：%T\n", str1, str1)
	// str2 := fmt.Sprintf("%.2f", f)
	// fmt.Printf("值：%v 类型：%T\n", str2, str2)
	// str3 := fmt.Sprintf("%t", t)
	// fmt.Printf("值：%v 类型：%T\n", str3, str3)
	// str4 := fmt.Sprintf("%c", b)
	// fmt.Printf("值：%v 类型：%T\n", str4, str4)

	//通过strconv把其他类型转换成string类型
	// var i int = 20
	// //参数1：int64的数值	参数2：传值int类型的进制
	// str1 := strconv.FormatInt(int64(i), 10)
	// fmt.Printf("值：%v 类型：%T\n", str1, str1)

	// var f float32 = 20.231313
	// str2 := strconv.FormatFloat(float64(f), 'f', 4, 64)
	// fmt.Printf("值：%v 类型：%T\n", str2, str2)

	//string类型转换成整型
	// str := "123456"
	// fmt.Printf("%v--%T\n", str, str)
	// //参数1：string数据	参数2：进制	参数3：位数 32 64 16
	// num, _ := strconv.ParseInt(str, 10, 64)
	// fmt.Printf("%v--%T\n", num, num)

	str := "123456"
	num, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("%v--%T\n", num, num)

}
