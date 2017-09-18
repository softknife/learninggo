/*
反射reflection
反射可以大大提供程序的灵活性, 使得interface{}有更大的发挥余地
反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
反射会将匿名字段作为独立字段(匿名字段本质)
想要利用反射修改对象状态,前提是interface.data是settable, 即pointer-interface
通过反射可以"动态"调用方法!
*/

package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
}

func (u user) hello() {

	fmt.Println("hello ", u.name)
}

func testReflection() {

}

func info(o interface{}) {

	/*
	要去反射是一个类型的值(这些值都实现了空interface)，
	首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，
	根据不同的情况调用不同的函数)。

	*/
	// type reflection 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(o)
	fmt.Println("type:", t.Name())

	// value reflection 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	v := reflect.ValueOf(o)
	fmt.Println("fields:")

	// field value reflection
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// method reflection
	for i := 0; i < t.NumMethod(); i++ {

		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)

	}

}
