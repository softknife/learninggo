/*
go语言中没有class,也就没有继承的概念, 只能嵌套
*/

package main

import "fmt"

// Human 人类基类
type Human struct {
	Sex int
}

type teacher struct {
	Human // 注意这里嵌套时,必须和被嵌套类型名字一样!!
	Name  string
	Age   int
}

func testStructEmbedded() {

	a := teacher{Name: "e", Age: 1, Human: Human{Sex: 0}}
	fmt.Println(a)
	a.Sex = 1 // 注意这里可以直接这样简写!!!
	// 但是为了避免和外部结构体中的成员重名,我们还可以如下写法
	a.Human.Sex = 0
	fmt.Println(a.Sex)
}
