package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：要统计1-120000的数字中哪些是素数？素数：只能被自身整除的为素数

1协程 统计1-30000
2协程 统计30001-60000
1协程 统计60001-90000
1协程 统计90001-120000
*/

var wg1 sync.WaitGroup

func test5(n int) {
	for num := (n-1)*30 + 1; num < n*30; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				//fmt.Println(num, "是素数")
			}
		}
	}
	wg1.Done()
}

func main() {
	star := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg1.Add(1)
		go test5(i)
	}
	wg1.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end - star)
}
