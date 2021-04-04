package main

import "fmt"

func func1() int {
	var a int
	fmt.Printf("当前数据",a)
	return a
}

func func2() int {
	var b int
	b = func1()
	fmt.Printf("当前数据",b)
	return b
}

func func3() int{
	var c int

	c = func2()
	fmt.Printf("当前数据",c)
	return c
}

func main()  {
	func1()
	func2()
	func3()
}