package main

import "fmt"

/*
结构体的字段类型可以是：基本数据类型、也可以是切片、Map以及结构体
如果结构体的字段类型是：指针、slice、和map的零值都是nil，即还没有分配空间
如果需要使用这样的字段，需要先make才能使用
*/
type Person3 struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

/*
结构体嵌套
*/
type User struct {
	Username string
	Password string
	Address  Address //表示User结构体嵌套Address结构体
	//Address //匿名结构体
}
type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {
	var p Person3
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "运动"
	p.Hobby[1] = "打游戏"
	p.Hobby[2] = "睡觉"

	p.map1 = make(map[string]string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "121212"

	fmt.Printf("%#v\n", p)
	fmt.Printf("%v", p.Hobby)

	var u User
	u.Username = "张三"
	u.Password = "2222"
	u.Address.Name = "哈哈"
	u.Address.Phone = "11111"
	u.Address.City = "北京"

	fmt.Printf("%#v", u)

}
