/*

 */

package main

import "fmt"

func testSiceUsage() {

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a的值: %v \n", a)

	// 使用slice访问数组
	s1 := a[3:10] // [4 5 6 7 8 9 10]
	s2 := a[4:]   // [5 6 7 8 9 10]

	fmt.Println("Slice使用", "\n s1:", s1, "\n s2:", s2)
	/*
			Slice使用
		 	s1: [4 5 6 7 8 9 10]
		 	s2: [5 6 7 8 9 10]
	*/

	// 通过slice修改数组
	s1[0] = 11
	fmt.Println("Slice修改数组\n a:", a, "\n s1:", s1, "\n s2:", s2)
	/*
			Slice修改数组
		 	a: [1 2 3 11 5 6 7 8 9 10]
		 	s1: [11 5 6 7 8 9 10]
		 	s2: [5 6 7 8 9 10]
	*/
}
