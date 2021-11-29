package main

import (
	"fmt"
	"log"
)

/*
	自定义error类型
*/
type PathError struct {
	Op   string
	Path string
	Err  error
}

/*
	实现error接口的Error方法
*/
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}
func main() {

	/*
		对于大多数函数，如果要返回错误
		可以将error作为多种返回值中的最后一个
	*/
	a, err := errorFunc()
	if err != nil {
		fmt.Println(a)
	}

	/*
		当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，
		但函数中 之前使用defer关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，
		并导致 逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止
		该函数接收任意类型的数据，比如整 型、字符串、对象等
	*/
	panic(404)
	//panic("network broken")

	/*
		recover()函数用于终止错误处理流程。
		一般情况下，recover()应该在一个使用defer 关键字的函数中执行以有效截取错误处理流程。
		如果没有在发生异常的goroutine中明确调用恢复 过程(使用recover关键字)，会导致该goroutine所属的进程打印异常信息后直接退出
	*/
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()

}
func foo() {

}
func errorFunc() (a int, err error) {
	return 1, err
}
