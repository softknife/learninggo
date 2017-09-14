/*
在某个包中定义的类型,在外部包中访问时,是无法直接绑定方法的,
但是,我们可以通过类型别名的方式来绑定方法!


本质上这就是一种语法糖，方法调用如下：
instance.method(args) -> (type).func(instance, args)

instance 就是Reciever.左边的称为 Method Value，右边则是 Method Expression,go推荐使用左边形式。
Method Value 是包装后的状态对象，总是与特定的对象实例关联在一起 (类似闭包，拐带私奔)，
而Method Expression 函数将 Receiver 作为第一个显式参数，调用时需额外传递。
二者本质上没有区别，只是Method Value 看起来更像面向对象的格式,且编译器会自动进行类型转换；
Method Expression更直观，更底层，编译器不会进行优化，会按照实际表达意义去执行，更易于理解。

*/

package main

import "fmt"

// TZ 是int的类型别名
type TZ int

func testTypealias() {

	var a TZ
	a.print()       // 这种写法叫做method value
	(*TZ).print(&a) // 这种写法叫做method expression

}

func (a *TZ) print() {
	fmt.Println("TZ")

}
