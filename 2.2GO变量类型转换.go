/*
go中不存在隐式类型转换,所有的类型转换必选显示声明
转换只能发生在两种相互兼容的类型之间
类型转换的格式:
	<ValueA> := <TypeOfValueA>(<ValueB>)
*/
package main

import "fmt"

func testTypeTransfrer() {

	// 相互兼容的两种类型之间才能进行转换
	var a float32 = 1.1
	b := int(a)

	fmt.Println(b)

}
