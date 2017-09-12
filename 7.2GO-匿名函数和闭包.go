/*
匿名函数不能作为最外层的函数
闭包
*/
package main

import (
	"fmt"
)

func testAnonamousFunc() {

	a := func() {
		fmt.Println("Anonamous Func A")
	}

	a()
}

func testClosure() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {

	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &y)
		return x + y
	}

}
