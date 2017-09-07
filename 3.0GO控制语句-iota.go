/*
iota枚举
Go里面有一个关键字iota,这个关键字用来声明enum的时候采用,默认值是0,每次调用加1
*/

package main

import "fmt"

const (
	B  float64 = 1 << (iota * 10)
	KB float64 = 1 << (iota * 10) // y= 10B
	MB                            // 常量声明省略值时,默认和之前一个值的字面相同, 这里隐士地说z= iota
	GB
)

const w = iota // 每次遇到一个const关键字,iota就会重置,此时w = 0

func testByteUnit() {

	fmt.Println(GB)
}
