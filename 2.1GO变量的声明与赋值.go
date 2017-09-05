/*
1.单个变量的声明与赋值

变量的声明格式: var <变量的名称> <变量的类型>
变量的赋值格式: <变量名称> = <表达式>
声明的同时赋值: var <变量名称> [变量类型] = <表达式>
*/

package main

import (
	"fmt"
)

var a int // 声明,注意这种声明的变量是全局的,意味着整个包中都可以访问使用

var b int = 123 // 声明的同时赋值,编译器建议我们忽略int,因为类型推导
var c = 123     // 可以简化为这样,类型推导

func variableDeclaration() {
	a = 1234 // 赋值

	d := 123 // 变量的声明和赋值最简写法,
	/*
		注意如果这样写,必须使用该变量,否则报错
		./2.1GO变量的声明与赋值.go:19:7: d declared and not used
		exit status 2
		Process exiting with code: 1
	*/

	fmt.Println(d)

}

/*
2.多个变量同时声明与赋值

	全局变量的声明可使用var()的方式进行简写
	全局变量的声明不可以省略var ,但可使用并行方式
	所有变量都可以使用类型推断
	局部变量不可以使用var()的方式简写,只能使用并行方式
*/
var (
	aaaa     = "hellow" // 常规方式定义
	bbb, ccc = 1, 2     // 并行方式定义变量
)

var abc, cde, fgq int // 多个变量同时声明

var k, j, i int = 2, 3, 4 // 多个变量声明的同时赋值
var y, u = 3, 44          // 省略写法,类型推导

func testMultiDeclaration() {

	abc, cde, fgq = 1, 3, 4 // 多个变量同时赋值
	h, l := 3, 45           // 最简写法

	fmt.Println(h, l)

}
