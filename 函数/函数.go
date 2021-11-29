package main

import "fmt"

func main() {

	/*
		函数
		小写字母开头的函数只在本包内可见，大写字母开头的函数才 能被其他包使用。
	*/
	b := add(1, 2)
	fmt.Println(b)

	fmt.Println(addArgs(1, 2, 3, 4, 5))

	fmt.Println(moreReturnVal(1, 2))
}

func add(x int, y int) int {
	return x + y
}

//不定参数 ...type实质是一个切片
func addArgs(args ...int) int {
	sums := 0
	for _, val := range args {
		sums += val
	}
	return sums
}

//任意类型的不定参数
func functionInterface(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		}
	}
}

//多返回值,如果不想使用某个返回值，可以用'_'替换
func moreReturnVal(x int, y int) (int, int) {
	x, y = y, x
	return x, y
}
