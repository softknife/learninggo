/*
Go没有异常机制, 但有panic/recover模式来处理错误
Panic可以在任何地方引发,但recover只有在defer调用的函数中有效
*/

package main

import "fmt"

func testPanic() {
	testPanic1()
	testPanic2()
	testPanic3()
}

func testPanic1() {
	fmt.Println("Func 1")
}

func testPanic2() {

	// 注意,这里的defer需要先写,不然不能把这个匿名函数注册到defer中!
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in func2")
		}
	}()

	panic("Panic in func2")

}

func testPanic3() {
	fmt.Println("Func 3")
}
