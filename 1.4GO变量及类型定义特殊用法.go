package main

import (
	"fmt"
)

// 1.

// 2.连续定义格式
/*
注意使用小括号
*/
// 变量连续定义
var (
	aaa   = 3.13
	name1 = "eric"
	name2 = "ingram"
)

// 常量连续定义
const (
	const1 = 1
	const2 = 2
)

// 类型定义
type (
	timeinterval float32
	gender       int
)

func testDef() {

	var bb timeinterval = 10.0

	fmt.Println(const1 + const2)

	fmt.Println(name1 + name2)

	fmt.Println("测试连续定义", bb)

}
