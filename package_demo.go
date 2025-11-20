package main

//在根目录中运行go mod init 项目名 会生成一个go.mod文件
import (
	"fmt"
	"golang/calc"
	T "golang/tools" //包别名
	//_ "my_module/tools" 匿名导入包
)

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)

	sub := calc.Sub(15, 3)
	fmt.Println(sub)

	sum1 := T.Sum(12, 3)
	println(sum1)
}
