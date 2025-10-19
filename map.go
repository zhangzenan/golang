package main

import "fmt"

func main() {

	//1、make创建map类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "男"
	// fmt.Println(userinfo)

	//2、map也支持在声明的时候填充元素
	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)

	//3、第三种创建map类型数据的方法
	// userinfo := map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)

	//for range循环遍历map类型的数据
	// for k, v := range userinfo {
	// 	fmt.Printf("key:%v  value:%v\n", k, v)
	// }

	//map类型数据的crud
	//1、创建 修改map类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// fmt.Println(userinfo)

	//创建修改map类型的数据
	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// 	"height":   "180cm",
	// }
	// userinfo["username"] = "李四"
	// fmt.Println(userinfo)

	// //获取 查找map类型的数据
	// v, ok := userinfo["age"]
	// fmt.Println(v, ok)

	// //删除map数据里面的key以及对应的值
	// delete(userinfo, "sex")
	// fmt.Println(userinfo)

	//元素为map类型的切片
	// var userinfo = make([]map[string]string, 3, 3)
	// fmt.Println(userinfo[0]) //map不初始化的默认值nil

	// if userinfo[0] == nil {
	// 	userinfo[0] = make(map[string]string)
	// 	userinfo[0]["username"] = "张三"
	// 	userinfo[0]["age"] = "20"
	// 	userinfo[0]["height"] = "180cm"
	// }
	// if userinfo[1] == nil {
	// 	userinfo[1] = make(map[string]string)
	// 	userinfo[1]["username"] = "张三1"
	// 	userinfo[1]["age"] = "21"
	// 	userinfo[1]["height"] = "181cm"
	// }
	// fmt.Println(userinfo)

	//map的值类型可以是切片
	// var userinfo = make(map[string][]string)
	// userinfo["hobby"] = []string{
	// 	"吃饭",
	// 	"睡觉",
	// 	"敲代码",
	// }
	// userinfo["work"] = []string{
	// 	"php",
	// 	"java",
	// 	"golang",
	// }
	// fmt.Println(userinfo)

	// for k, v := range userinfo {
	// 	//fmt.Printf("key:%v value:%v\n", k, v)
	// 	for _, value := range v {
	// 		fmt.Printf("key:%v value:%v\n", k, value)
	// 	}

	// }

	//map类型也是引用数据类型
	var userinfo1 = make(map[string]string)
	userinfo1["username"] = "张三"
	userinfo1["age"] = "20"
	userinfo2 := userinfo1

	userinfo2["username"] = "李四"
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)

}
