package main

import "fmt"

type Integer int

/*
	(a *Integer) 是对象中的属性 (b Integer) b是传进来的参数
	注意:
		如果想对属性的值进行修改，就需要传入指针，因为在Go中所有的类型均为值类型，不会改变内存中的值
		对于切片，map，channel，接口来说看起来像引用类型
*/
func (a *Integer) Add(b Integer) {
	*a += b
}

/*
	结构体:可以类比Class
*/
type Rect struct {
	//定义属性
	x, y          float64
	width, height float64
}

/*
	定义成员方法
*/
func (r *Rect) Area() float64 {
	return r.width * r.height
}

/*
	定义构造方法
		在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以 NewXXX来命名，表示“构造函数”
*/
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main() {
	var a Integer = 1
	a.Add(3)
	fmt.Println(a)

	/*
		结构体初始化方式
		注意:
			未进行显式初始化的变量都会被初始化为该类型的零值
			例如bool类型的零 值为false，int类型的零值为0，string类型的零值为空字符串
	*/
	/*rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}*/

	rec1 := NewRect(1, 2, 3, 4)
	area := rec1.Area()
	fmt.Println(area)
}
