package main

import "fmt"

func main() {

	/*
		if-else
	*/
	a := 3
	if a < 5 {
		fmt.Println(11)
	} else {
		fmt.Println(22)
	}

	/*
		switch
		Go语言不需要用break来明确退出一个case
		只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
	*/
	i := 2
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}

	fmt.Println("--------------------")
	/*
		for 循环
	*/
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
