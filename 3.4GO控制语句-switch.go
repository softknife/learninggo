/*
选择语句switch
	可以使用任何类型或表达式作为条件语句
	不需要写break, 一旦条件符合自动终止
	如希望继续执行下一个case ,需要使用fallthrough语句
	支持一个初始化表达式(可以是并行方式),右侧需跟分号
	左大括号必须和条件语句在同一行
*/

package main

import "fmt"

func testSwitch() {

	// 形式1-常规用法
	a := 1
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	}

	fmt.Println(a)

	// 形式2-表达式作为case
	b := 1
	switch {
	case b >= 0:
		fmt.Println("b >= 0")
	case b >= 1:
		fmt.Println("b >= 1")
	}

	fmt.Println(b)

	// 形式3-并行方式初始化一个表达式
	switch a := 1; {
	case a >= 0:
		fmt.Println(" a >= 0")
	case a >= 1:
		fmt.Println("a >= 1")
	}
}
