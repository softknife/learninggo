/*
并发concurrency
	很多人都是冲着Go大肆宣扬的高并发而来, 但其实从源码解析来看, goroutine只是由官方实现的超级"线程池"而已,
不过话说回来,每个实例4-5kb的栈内存占用和由于实现机制而大幅减少的创建和销毁开销,是制造Go号称的高并发的根本原因.
另外,goroutine的简单易用,也在语言层面上给予了开发者巨大的便利.
	并发不是并行 : concurrency is not parallelism
	并发主要由切换时间片来实现"同时"运行, 而并行则是直接利用多核实现多线程的运行, 但Go可以设置使用核数,以发挥
多核计算机的能力.
goroutine奉行通过通信来共享内存, 而不是共享内存来通信. ???


Channel:
	channel是goroutine沟通的桥梁,大都是阻塞同步的
	通过make创建,close关闭
	channel是引用类型
	可以使用for range来迭代不断操作channel
	可以设置单向或者双向通道
	可以设置缓存大小, 在未被填满前不会发生阻塞

Select
	可以处理一个或者多个channel的发送或者接收
	同时有多个可用的channel时,按随机顺序处理
	可用空的select来阻塞main函数
	可设置超时

*/

package main

import (
	"fmt"
)

func testGoRoutine() {

	/// 1.一般使用
	// 使用make创建channel,channel内部绑定类型是布尔类型
	c := make(chan bool)

	// 使用goroutine执行一段任务代码块
	go func() {
		fmt.Println("Go !!")
		c <- true // 写入
	}()

	<-c // 读出
	/// 上述channel是无缓存的, 同步执行!

	/// 2.for range
	c1 := make(chan bool)
	go func() {
		fmt.Println("go go !!")
		c1 <- true
		close(c1)
	}()

	for v := range c1 {
		fmt.Println(v)
	}

}
