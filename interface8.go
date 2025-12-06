package main

import "fmt"

//Golang中空接口和类型断言使用细节

type Address1 struct {
	Name  string
	Phone int
}

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])

	//fmt.Println(userinfo["hobby"][1])//cannot index userinfo["hobby"] (map index expression of type interface{})

	var address = Address1{
		Name:  "李四",
		Phone: 12344445,
	}
	fmt.Println(address.Name)

	userinfo["address"] = address
	//fmt.Println(userinfo["address"].Name)// 报错

	//断言userinfo["hobby"]是否是切片类型
	hobby2, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[1])

	address2, _ := userinfo["address"].(Address1)
	fmt.Println(address2.Name, address2.Phone)

}
