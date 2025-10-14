package main

import "fmt"

func main() {

	//注意：++(自增)和--（自减）在Go语言中是单独的语句，并不是运算符。
	//注意：在golang中，++和--只能独立使用，不能作赋值运算

	// var i int = 8
	// var a int
	//a=i++//错误，i++只能独立使用
	//a=i--//错误，i--只能独立使用
	//++和--只能放在变量的后面不能放在前面
	// fmt.Println(i, a)

	//go语言中是没有while语句的，我们可以通过for代替
	// for {
	// 	循环语句
	// }

	// var str = "你好golang"
	// for k, v := range str {
	// 	fmt.Printf("key=%v  val=%c\n", k, v)
	// }

	//var arr = []string{"php", "java", "node", "golang"}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for k, v := range arr {
	// 	fmt.Println(k, v)
	// }

	// for _, val := range arr {
	// 	fmt.Println(val)
	// }

	// var extname = ".html"
	// switch extname {
	// case ".html":
	// 	fmt.Println("txt/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")

	// }

	// switch extname := ".html"; extname {
	// case ".html":
	// 	fmt.Println("txt/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	// var n = 8
	// switch n {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// 	break //golang中的break可以写也可以不写
	// case 2, 4, 6, 8, 10:
	// 	fmt.Println("偶数")
	// 	break
	// }

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	//fallthrough语法可以执行满足条件的case的下一个case,是为了兼容C语言中的case设计，只能穿透一层
	// var age = 30
	// switch {
	// case age < 24:
	// 	fmt.Println("好好学习")
	// case age >= 24 && age <= 60:
	// 	fmt.Println("好好赚钱")
	// 	fallthrough
	// case age > 60:
	// 	fmt.Println("注意身体")
	// default:
	// 	fmt.Println("输入错误")
	// }

	//break 跳出当前循环
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	//在多重循环中，可以用标号label标出想跳出的循环
	// label:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 3 {
	// 				break label
	// 			}
	// 			fmt.Println(j)
	// 		}
	// 	}

	//continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用
	// for i := 1; i <= 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	//在continue语句后添加标签时，表示开始标签对应的循环
	// label2:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 3 {
	// 				continue label2
	// 			}
	// 			fmt.Printf("i=%v j=%v\n", j, j)
	// 		}
	// 	}

	//goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定帮助
	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto lable3
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
lable3:
	fmt.Println("ccc")
	fmt.Println("ddd")

}
