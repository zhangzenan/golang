package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %d:%d:%d\n", year, month, day, hour, minute, second)

	//go语言的诞生日期是2006年01月02日15点04分（记忆口诀为2006 1 2 3 4）
	/**
	2006	年
	01	月
	02	日
	03	时 12小时制 写成15 表示24小时制
	04	分
	05	秒
	**/
	//2025-02-03 01:23:23
	str := timeObj.Format("2006-01-02 03:04:05")
	fmt.Println(str)

	//2025-02-03 01:23:23
	str2 := timeObj.Format("2006/01/02 03:04:05")
	fmt.Println(str2)

	str3 := timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(str3)

	//获取当前时间戳
	unixtine := timeObj.Unix()
	fmt.Println(unixtine)

	//获取当前纳秒时间戳
	unixNatime := timeObj.UnixNano()
	fmt.Println(unixNatime)

	//时间戳转换成日期字符串
	unixTime := 1762264232
	timeObj1 := time.Unix(int64(unixTime), 0)
	var str4 = timeObj1.Format("2006-01-02 15:04:05")
	fmt.Println(str4)

	//日期字符串转换成时间戳
	var str5 = "2025-11-04 21:50:32"
	var tmp = "2006-01-02 15:04:05"
	timeObj2, _ := time.ParseInLocation(tmp, str5, time.Local)
	fmt.Println(timeObj2)
	fmt.Println(timeObj2.Unix())

	fmt.Println(time.Millisecond) //1毫秒
	fmt.Println(time.Second)      //1秒

	//时间操作函数
	fmt.Println(timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println(timeObj)

	//定时器
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop() //终止这个定时器继续执行
			break
		}
	}

	//sleep方式实现定时器
	for {
		time.Sleep(time.Second)
		fmt.Println("我在定时执行任务")
	}

}
