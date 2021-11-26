package main

import (
	"fmt"
)

func main() {
	/**
	基本数据类型			数据范围
	布尔 bool
	整型{
		int8			-128 ~ 127
		byte			0-255
		int16			-32768~32767
		int				平台相关
		int32			-2147483648~2147483647
		int64			-9223372036854775808~9223372036854775807
		uint			平台相关
		uintptr
	}
	浮点型{
		float32
		float64
	}
	字符串 string
	字符 rune
	错误类型 error

	复合数据类型
	指针，数组，切片，字典，通道，结构体，接口
	*/

	/**
	注意:
		1.不同数据类型之间的变量不能进行相互比较
		2.浮点数不能使用 == 进行比较
	*/
	var a int32 = 10
	var b int64 = 10
	fmt.Println(a)
	fmt.Println(b)

	str := "Hello"
	/**
	获取字符串第i个字符 str[i],
	获取字符串长度 len(str)
	字符串初始化后不能被修改
	*/
	c := str[1]
	fmt.Println(c)
	fmt.Println(len(str))

	/**
	数组定义
	*/
	arr := [10]int{1, 2, 3}
	fmt.Println(arr[4])

	/**
	元素访问
	*/
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}

	for key, val := range arr {
		fmt.Println("key = ", key, ",val", val)
	}

	/**
	注意
		数组是一个值类型，所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作，
		如果将数组作为函数的参数类型，则在函数调用时该 参数将发生数据复制。
		因此，在函数体中无法修改传入的数组的内容，因为函数内操作的只是所 传入数组的一个副本。
		如果想要修改，则需要以引用的方式进行值传递
	*/
	arrByPtrTransform(&arr)
	fmt.Println(arr[1])
	arrByValTransform(arr)
	fmt.Println(arr[1])
}
func arrByPtrTransform(arr *[10]int) {
	arr[1] = 4
}

func arrByValTransform(arr [10]int) {
	arr[1] = 4
}
