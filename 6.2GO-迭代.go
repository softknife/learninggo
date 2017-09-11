/*
注意,for range循环时, 是对集合中元素的拷贝,所以,直接对循环出的变量进行操作是不可行的
*/

package main

import "fmt"

func testForRange() {

	sm := make([]map[int]string, 5)
	for _, v := range sm {

		// 这里v是拷贝出来的值!
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)

	}

	fmt.Println(sm)
	/*
		map[1:ok]
		map[1:ok]
		map[1:ok]
		map[1:ok]
		map[1:ok]
		[map[] map[] map[] map[] map[]]
	*/

	// 如下写法才是OK的!
	// 注意,对集合进行for range时,不能省略value,可以省略key
	for i := range sm {

		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}

	fmt.Println(sm)
	/*
		map[1:OK]
		map[1:OK]
		map[1:OK]
		map[1:OK]
		map[1:OK]
		[map[1:OK] map[1:OK] map[1:OK] map[1:OK] map[1:OK]]
	*/

}
