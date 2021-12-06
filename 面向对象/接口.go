package main

import "fmt"

/*
	接⼝口是⼀一个或多个⽅方法签名的集合，任何类型的⽅方法集中只要拥有与之对应的全部⽅方法，
	就表⽰示它 "实现" 了该接⼝，⽆无须在该类型上显式添加接⼝声明
	接⼝命名习惯以 er 结尾，结构体。
	接⼝只有⽅方法签名，没有实现。
	接⼝没有数据字段。
	可在接⼝中嵌⼊入其他接口。
	类型可实现多个接⼝
*/
type Stringer interface {
	String() string
}
type Printer interface {
	Stringer //接口嵌套
	Print()
}
type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d,%s", self.id, self.name)

}
func (seft *User) Print() {
	fmt.Println(seft.String())
}
func main() {
	var t Printer = &User{1, "Tom"}
	t.Print()
}
