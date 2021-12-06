package main

import (
	"fmt"
	"time"
)

/*
	使用channel在两个或 多个goroutine之间传递消息。channel是进程内的通信方式
	channel是类型相关的 一个channel只能传递一种类型的值 这个类型需要在声 明channel时指定

*/
func main() {
	//声明一个传递类型为int的channe
	//var chanName chan int

	//定义一个channel
	ch := make(chan int)

	/*
		将一个数据写入channel
	*/
	ch <- 1

	/*
		向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据
		如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel 中被写入数据为止
	*/
	value := <-ch
	fmt.Println(value)

	/*
		select 处理异步IO问题,一旦其中一个文件句柄发生了IO动作，该select()调用就会被返回
		每个 case语句都必须是一个面向channel的操作
	*/

	//实现了一个随机向ch中写入一个0或者1 的过程。当然，这是个死循环
	ch2 := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
	fmt.Println(ch2)

	/*
		创建一个带缓冲的channel,创建了一个大小 为1024的int类型channel
		即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被 填完之前都不会阻塞
	*/
	c := make(chan int, 1024)
	//从channel中读取数据
	for i := range c {
		fmt.Println("Received:", i)
	}

	/*
		使用select实现超时机制
	*/

	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
		// 从ch中读取到数据
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
	}

	/*
		单向channel的定义
	*/

	//var ch1 chan int // ch1是一个正常的channel，不是单向的
	//var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
	//var ch3 <-chan int // ch3是单向channel，只用于读取int数据

	//ch4 := make(chan int)
	//ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
	//ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel

	/*
		关闭channel
		x, ok := <-ch ok为false代表channel已经关闭
	*/
	close(ch)
}
