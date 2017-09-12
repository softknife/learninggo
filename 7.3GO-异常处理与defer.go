/*
defer的执行方式类似其他语言中的析构函数, 在函数体执行结束后,按照调用顺序的相反顺序逐个执行
即使函数发生严重错误也会执行!!!
支持匿名函数的调用
常用于资源清理,文件关闭,解锁以及记录时间等操作
通过与匿名函数配合可在return之后修改函数计算结果
如果函数体内某个变量作为defer时匿名函数的参数, 则在定义defer时即已经获得了拷贝,否则则是引用某个变量的地址


*/

package main

import "fmt"

func testDefer() {

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	// 输出的结果2,1,0 , 这里i作为参数传入,每次都是对i的拷贝
}

func testDeferAnonamous() {

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	// 结果: 3,3,3 因为这里是引用i变量!!!
}
