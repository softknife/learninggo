/*
Reslice
	Reslice时索引以被slice的切片为准
	索引不可以超过被slice的切片的容量cap()值
	索引越界不会导致底层数组的重新分配,而是引发错误

Append
	可以在slice尾部追加元素
	可以将一个slice追加在另一个slice尾部
	如果最终长度未超过追加到slice的容量,则返回原始slice
	如果超过追加到的slice的容量, 则将重新分配数组,并拷贝原始数据

Copy
	copy(x,y)表示将y内容到x中
*/

package main

import (
	"fmt"
)

func testResliceAppendCopy() {

	// Reslice
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}

	sa := a[2:6]
	fmt.Println(len(sa), cap(sa))

	sb := sa[1:3]
	// sb := sa[6:7] 这样越界,会报错
	fmt.Println("sa:", string(sa), "\nsb:", string(sb))
	/*
		sa: cdef
		sb: de
	*/

	// Append
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	/*
		0xc420012060
		[0 0 0 1 2 3] 0xc420012060
	*/

	// Copy
	s3 := []int{1, 2, 3, 4, 5, 6}
	s4 := []int{7, 8, 9}
	copy(s4, s3)
	fmt.Println(s4) // [1 2 3] 超出范围的值直接被忽略

	s5 := make([]int, 4)
	copy(s5, s3[2:5])
	fmt.Println(s5) // [3 4 5 0]

}
