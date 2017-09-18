/*
如果想通过反射去修改字段, 那么该字段必须是可读写的
*/

package main

import (
	"fmt"
	"reflect"
)

// User 测试reflect修改属性
type User struct {
	Name string
}

func testReflectModifyField() {

	x := 12334
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(1234456)

	fmt.Println(x)

	// 组合类型 使用反射修改!
	u := User{Name: "aaa"}
	set(&u)
	fmt.Println(u)
}

func set(o interface{}) {

	v := reflect.ValueOf(o)

	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("不可修改")
		return
	}

	v = v.Elem()

	//f := v.FieldByName("name") // 注意这里小写的字段不能通过reflect修改,即使在一个包中
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if !f.CanSet() {

		fmt.Println("该字段不能修改")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("byebye")
	}

}
