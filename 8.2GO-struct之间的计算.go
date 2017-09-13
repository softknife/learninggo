/*
不同结构体类型,即使成员变量完全相同,也不能在一起计算
*/

package main

import "fmt"

type person2 struct {
	Name string
	Age  int
}

func testStructCompare() {

	a := person{Name: "nilson", Age: 5}
	// b := person2{Name: "nilson", Age: 5}

	// fmt.Println(a == b)
	// invalid operation : a == b (mismatched types person and person2)

	c := person{Name: "eric-ingram", Age: 10}
	fmt.Println(a == c)
}
