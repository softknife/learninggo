/*
判断语句if
	条件表达式没有括号
	支持一个初始化表达式(可以是并行方式)
	左大括号必须和条件语句或else在同一行
	支持单行模式
	初始化语句中的变量为block级别, 同时隐藏外部同名变量
	1.0.3版本中的编译器BUG
*/

package main

import "fmt"

func testIF() {

	a := true
	if a, b, c := 1, 2, 3; a+b+c > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("小于等于6")
		fmt.Println(a) // 这里a是内部的
	}

	fmt.Println(a) // 这里的a是外部的
}
