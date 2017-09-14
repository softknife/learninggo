/*
某个类型变量的绑定方法
和
某个类型指针变量的绑定方法不同
变量的绑定方法中操作变量不影响外部, 但是指针绑定的方法内部操作变量时会影响指针指向的变量


instance.method(args)：可以用 value 或 pointer 变量调用所有绑定的任何方法，编译器会自动进行类型转换，转换后的效果和methods的定义语义一致，和调用者的样式没有关系。
比如原method定义的Receiver是value形的，即使使用pointer 实例变量调用该方法，也是value值拷贝；
比如原method定义的Receiver是pointer形的，即使使用value 实例变量调用该方法，也能改变Receiver实例状态。
*/

package main

import "fmt"

type valueMethod struct {
	name string
}

type pointerMethod struct {
	name string
}

func testValuePoint() {

	a := pointerMethod{name: "原始值a"}
	a.print()
	fmt.Println(a.name)

	a1 := &pointerMethod{name: "定义指针a"}
	a1.print()
	fmt.Println(a1.name)
	// a和a1调用print方法时,作用一样,因为编译器会帮我们做处理,处理后与print()方法定义的语义一致!

	b := valueMethod{name: "原始值b"}
	b.print()
	fmt.Println(b.name)

	// AAA
	// AAA
	// BBB
	// 原始值b

}

func (a *pointerMethod) print() {

	a.name = "AAA"
	fmt.Println(a.name)
}

func (b valueMethod) print() {
	b.name = "BBB"
	fmt.Println(b.name)
}
