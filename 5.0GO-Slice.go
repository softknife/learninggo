/*
slice特性:
	其本身并不是数组,它指向底层的数组
	作为变长数组的替代方案, 可以关联底层数组的局部或全部
	为引用类型
	可以直接创建或者从底层数组获取生产
	使用len()获取元素个数,cap()获取容量
	一般使用make()创建
	如果多个slice指向相同底层数组,其中一个的值改变会影响全部
	创建方法:make([]T,len,cap)
		- cap可以省略,默认和len值相同
		- len表示存数的元素个数,cap表示容量
*/

package main

import "fmt"

func testSice() {

	// 定义一个slice,注意和数组的区别
	var s1 []int
	fmt.Println(s1)

	// 创建一个切片
	s2 := make([]int, 2, 10)
	fmt.Println(len(s2), cap(s2))

}
