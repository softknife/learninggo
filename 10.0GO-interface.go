/*
接口是一个或者多个方法签名的集合
只要某个类型拥有该接口的所有方法签名,即算实现该接口,无需显式声明实现了哪个接口, 这称为Structural Typing
	所以,按照上面这种说法, 空接口某人被所有类型实现!! interface{}可以被看做一个基类

接口只有方法声明, 没有实现,没有数据字段
接口可以匿名嵌入其他接口,或者嵌入到结构中
将对象赋值给接口时, 会发生拷贝,而接口内部存储的是指向这个复制品的指针,既无法修改复制品的状态,也无法获取指针
只有当接口存储的类型和对象都为nil时,接口才等于nil
接口调用不会做receiver的自动转换
接口同样支持匿名字段方法
接口也可实现类型OOP中的多态
空接口可以作为任何类型数据的容器

*/

package main

import "fmt"

type usb interface {
	name() string
	connecter
}

type connecter interface {
	connect()
}

type phoneConnecter struct {
	Name string
}

func (pc phoneConnecter) name() string {
	return pc.Name
}

func (pc phoneConnecter) connect() {
	fmt.Println("make connect:", pc.Name)
}

func testInterface() {
	a := phoneConnecter{Name: "iPhone x"}
	a.connect()

	b := &phoneConnecter{Name: "iPhone 8"}
	b.connect()

	disconnect(a)

	// make connect: iPhone x
	// make connect: iPhone 8
	// disconnect: iPhone x

	disconnect1(b)
	// 由于disconnect1方法接收的参数是接口类型, 那么编译器不会自动将b转换为值类型传递进去
	// 此时传递进去的类型是 *main.phoneConnecter类型的值!!
}

func disconnect(mc usb) {
	if pc, ok := mc.(phoneConnecter); ok { //OK pattern 类型转换!!
		fmt.Println("disconnect:", pc.Name)
		return
	}
}

// 我们还可以声明一个抽象的接口,在函数内部判断具体类型!
func disconnect1(usb interface{}) {

	switch v := usb.(type) { // 这种写法学名叫:type switch
	case phoneConnecter:
		fmt.Println("disconnect1:", v.Name)
	default:
		fmt.Println("Unknown device")
	}
}
