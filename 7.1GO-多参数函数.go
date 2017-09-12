/*
不定长参数
变参必须放到最后
*/

package main

import "fmt"

func testInvoke() {

	testChangeParameter("test", 1, 2, 3)
}

func testChangeParameter(a string, b ...int) {

	fmt.Println(a, b)
}
