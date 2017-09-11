/**
类似其他语言中的哈希表或者字典, 以key-value形式存储数据
key必须是支持==或者!=比较运算的类型, 不可以是函数,map或者slice
Map查找比线性搜索快很多, 但比使用索引访问数据的类型慢100倍
Map使用make()创建, 支持:=这种简写方式

make([keyType]valueType,cap) , cap表示容量, 可省略
超出容量时会自动扩容, 但尽量提供一个合理的初始值
使用len()获取元素个数

键值对不存在时自动添加, 使用delete()删除某键值对
使用for range 对map和sice进行迭代操作
*/

package main

import "fmt"

func testMap() {

	m := make(map[int]string) // 不声明容量
	m[1] = "OK"
	fmt.Println(m)

	delete(m, 1)

	a := m[1]

	fmt.Println(a)

}
