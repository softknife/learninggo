/*
一把我们是通过slice完成该操作
*/
package main

import (
	"fmt"
	"reflect"
)

type manager struct {
	title string
	user
	size int
}

func testAnonamousFields() {

	m := manager{user: user{name: "eric"}, title: "测试匿名字段"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))
	// reflect.StructField{Name:"user", PkgPath:"main", Type:(*reflect.rtype)(0x10da9a0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}

}
