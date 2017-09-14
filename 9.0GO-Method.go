/*
GO中虽然没有class,但是依旧有method
 通过显式说明receiver,来实现与某个类型的组合
 只能为同一个包中的类型定义方法!!!
 Receiver可以是类型的值或者指针
 不存在方法重载
 可以使用值或指针来调用方法,编译器会自动完成转换
 从某种意义上来说,方法是函数的语法糖, 因为receiver其实就是方法所接收的第一个参数(Method value VS Mehtod Expression)
 如果外部结构和嵌入结构存在同名方法, 则优先调用外部结构的方法
 类型别名不会拥有底层类型所附带的方法!!!! 但是可以再从新绑定自己的方法!
 方法可以调用结构中的非公开字段
*/

package main

import "fmt"

type methodStruct struct {
	Name string
}

func testMethod() {

	a := methodStruct{}
	a.print()
	// MethodStruct 1 {}

	b := methodStruct1{}
	b.print()
}

func (a methodStruct) print() {
	fmt.Println("MethodStruct 1", a)
}

// 不同类型的变量可以绑定相同名字的方法!
type methodStruct1 struct {
	Name string
}

func (a1 methodStruct1) print() {
	fmt.Println("methodstruct 2", a1)
}
