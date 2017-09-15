/*
大集合(超集)接口可以转换为子集接口,
这与
派生类(子类)与超类(父类)的关系正好相反
*/
package main

import (
	"fmt"
)

type engine interface {
	cylinder() int
}

type car interface {
	run()
	engine
}

type lamborghini struct {
	name string
}

func (car lamborghini) run() {

	fmt.Println(car.name, "跑起来!")

}

func (car lamborghini) cylinder() int {

	return 4
}

func testSuperCollectionInterface() {

	car := lamborghini{name: "兰博基尼"} // 这里lamborghini实现了car接口中定义的所有方法
	var e engine                     // engine属于car的子集
	e = engine(car)                  // 可见,可以将超集强转为子集; 但是相反转换是不可以的

	fmt.Println(e.cylinder())
}
