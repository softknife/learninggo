/*
通过反射调用方法
*/

package main

import (
	"fmt"
	"reflect"
)

// SayHelloTo 测试reflect 调用方法
func (u User) SayHelloTo(name string, lastName string) {
	fmt.Println("hello", name, lastName, ",my name is ", u.Name)
}

func testReflectCallMethod() {

	u := User{Name: "小明"}
	v := reflect.ValueOf(u)

	// 取出方法
	mv := v.MethodByName("SayHelloTo")

	// 定义参数,使用slice定义参数
	args := []reflect.Value{reflect.ValueOf("eric"), reflect.ValueOf("Ingram")}

	// 调用!
	mv.Call(args)

}

/*
上述测试,我们可以发现只有对包外可见等级的属性或者方法才能通过reflect修改或者调用!!!❤️
*/
