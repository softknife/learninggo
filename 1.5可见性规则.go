/*
在go语言中,使用大小写来决定该
常量,变量,类型,接口,结构或者函数
是否可以被外部包所调用
根据约定,函数名首字母小写即为private
*/

package main

import (
	"fmt"
)

func getNumber() bool {
	return true
}

/*
函数名首字母大写即为public
*/
func MyPrint() {
	fmt.Println("测试对包外可见的函数")
}
