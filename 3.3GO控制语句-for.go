/*
循环语句for
	Go只有for一个循环语句关键字,但是支持3中形式
	初始化和步进表达式可以是多个值
	条件语句每次循环都会被重新检查,因此不建议在条件语句中使用函数,尽量提前计算好条件并以变量或者常量代替
	左大括号必须和条件语句在同一行
*/

package main

import "fmt"

func testForLoop() {

	a := 1
	// for形式1
	for {
		a++
		if a > 1 {
			break
		}

		fmt.Println("for形式1", a)
	}

	// for形式2
	a = 1
	for a < 3 {
		a++
	}
	fmt.Println("for形式2", a)

	// for形式3
	a = 1
	for i := 0; i < 3; i++ {
		a++
	}

	fmt.Println("for形式3", a)

}
