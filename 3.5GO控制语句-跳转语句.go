/*
跳转语句goto,break,continue
	三个语法都可以配合标签使用
		标签和哪个for同一级别,就跳转到哪个for执行该行为
	标签名区分大小写, 若不适用会造成编译错误
	Break与continue配合标签可用于多层循环的跳出
	Goto是调整执行位置, 与其他2个语句配合
*/

package main

import "fmt"

func testLABEL() {

	// 外部为无限循环-break
LABEL:
	for {
		for i := 0; i < 10; i++ {

			if i > 3 {
				break LABEL // 由于外部有一个无限循环,所以这里需要指出具体break位置
			} else {
				fmt.Println("外部for循环,使用break控制", i)
			}
		}

	}

	// 内部为无限循环-continue/break
LABEL1:
	for i := 0; i < 10; i++ {

		for {
			fmt.Println("内部for循环,使用continue控制", i)
			continue LABEL1
		}
	}

	// LABEL3: 这样写明显是错误的
	// for {
	// 	for i := 0; i < 10; i++ {
	// 		if i > 3 {
	// 			continue LABEL3
	// 		}
	// 	}
	// }

	// 等同于
	for i := 0; i < 10; i++ {
		for {
			fmt.Println("break默认行为", i)
			break
		}
	}

	// 外部无限循环-goto使用
	// LABEL2:
	// for { 这样写是有问题的,外部无限循环会循环死
	// 	for i := 0; i < 10; i++ {
	// 		if i > 3 {
	// 			goto LABEL2
	// 		}
	// 	}
	// }

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL2
			}
		}
	}
LABEL2:

	fmt.Println("结束")

}
