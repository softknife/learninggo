/*
布尔型:bool
	长度:1字节
	取值范围:true,false
	注意事项:不可以用数字代表true或者false

整型:int/uint
	根据运行平台可能为32或64位

8为整型:int8/uint8
	长度:1字节
	取值范围:-128~127/0~255

字节型:byte(uint8别名)

16位整型:int16/uint16
	长度:2字节
	取值范围: -32768~32767/0~65535

32位整型:int32(rune)/uint32
	长度:4个字节
	取值范围:-2^32/2~2^32/2-1 / 0~2^32-1

64位整型:int64/uint64
	长度:8字节
	取值范围:-2^64/2~2^64/2-1 / 0~2^64-1

浮点型:float32/float64
	长度:4/8字节
	小数位:精确到7/15小数位

复数:complex64/complex128
	长度:8/16字节

足够保存指针的32位或者64位整数型:uintptr

其他值类型:
	array,struct,string

引用类型:
	slice,map,chan

接口类型:interface

函数类型:func

*/

/*
类型零值
	零值并不等于空值, 而是当变量被声明为某种类型后的默认值,
	通常情况下值类型的默认值是0,bool为false,string为空字符串
*/
package main
