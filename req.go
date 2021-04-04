package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("这是一个滞后函数！")

	a := 1

	if a > 10{
		fmt.Printf("当前数据不队")
	}else{
		fmt.Println("当前数据是：", a)
	}
	time.Sleep(2* time.Second)
}
