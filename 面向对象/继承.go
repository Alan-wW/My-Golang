package main

/*
	定义一个父类
*/
type Base struct {
	Name string
}

/*
	父类的两个方法
*/
func (base *Base) Foo() {

}
func (base *Base) Bar() {

}

/*
	定义一个子类，子类继承Base类
*/
type Foo struct {
	Base
	//在Go语言中，还可以以指针方式从一个类型“派生”
	//*Base
}

/*
	子类重写了父类的Bar方法
*/
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

/*
	接口组合中的名字冲突问题
		所有的Y类型的Name成员的访问都只会访问到最外层的那个Name变量，X.Name变量相当于被隐藏起 来了
*/
type X struct {
	Name string
}
type Y struct {
	X
	Name string
}

/*
	Go语言对关键字的增加非常吝啬，其中没有private、protected、public这样的关键 字。
	要使某个符号对其他包(package)可见(即可以访问)，需要将该符号定义为以大写字母 开头
*/
