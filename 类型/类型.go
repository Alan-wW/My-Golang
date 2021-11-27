package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

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

	/*
		切片的创建
	*/

	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]                              // 取数组前五个元素
	fmt.Println("Elements of myArray: ")
	for _, val := range mySlice {
		fmt.Println(val)
	}

	//直接创建一个初始长度为5的切片
	//mySlice2 := make([]int,5)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间:
	//mySlice3 := make([]int, 5, 10)

	//直接创建并初始化包含5个元素的数组切片:
	mySlice4 := []int{1, 2, 3, 4, 5}
	mySlice5 := []int{1, 3}

	//如果需要往向mySlice4已包含的5个元素后面继续新增元素，
	//可以使用append()函数。 下面的代码可以从尾端给mySlice加上3个元素，从而生成一个新的数组切片:
	mySlice = append(mySlice4, 1, 2, 3)
	fmt.Println("before", mySlice)

	//向mySlice4中追加mySlice5中的元素
	mySlice = append(mySlice, mySlice5...)
	fmt.Println("after", mySlice)

	//元素复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)

	/*
		map 的创建
	*/

	//var personDB map[string] PersonInfo
	//personDB = make(map[string] PersonInfo)
	//personDB = map[string]PersonInfo{
	//	"1234": PersonInfo{"1", "Jack", "Room 101,..."},
	//}
	personDB := make(map[string]PersonInfo)
	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	// 从这个map查找键为"1234"的信息
	person, ok := personDB["1234"]
	if ok {
		fmt.Println(person)
	} else {
		fmt.Println("Not exist 1234")
	}

	/*
		map 元素删除
	*/
	delete(personDB, "1")
	fmt.Println(personDB)
}
func arrByPtrTransform(arr *[10]int) {
	arr[1] = 4
}

func arrByValTransform(arr [10]int) {
	arr[1] = 4
}
