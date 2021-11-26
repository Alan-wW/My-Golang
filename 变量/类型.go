package main

import "fmt"

func main() {
	/**
	变量声明
	*/

	/*var v1 int
	var v2 string
	var v3 [10]int //数组
	var v4 []int //数组切片
	var v5 *int //指针
	var v6 map[string]int //字典
	var v7 struct{ // 结构体
		f int
	}*/

	/**
	多变量声明
	*/

	/*var(
		v8 int
		v9 int
	)*/

	/**
	变量初始化
	*/

	var a int = 10
	//var b = 10
	c := 10

	/**
	变量交换
	*/
	a, c = c, a

	/**
	匿名变量,如果不想使用某个变量可以用 "_"代替
	*/
	_, _, d := getName()
	fmt.Println(d)

	/**
	常量
	*/
	const v1 = 10

	/**
	枚举
	*/
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Saturday)
}
func getName() (a int, b int, c int) {
	return 1, 2, 3
}
