/*
## Go 标记
Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。
如以下 GO 语句由 6 个标记组成：

fmt.Println("Hello, World!")

6 个标记是(每行一个)：
1. fmt
2. .
3. Println
4. (
5. "Hello, World!"
6. )



## 行分隔符
在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，
因为这些工作都将由 Go 编译器自动完成。
如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，
但在实际开发中我们并不鼓励这种做法。



## 标识符 同c标识符规则一致
标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。



## 关键字
下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
break		default			func		interface		select
case		defer			go			map				struct
chan		else			goto		package			switch
const		fallthrough		if			range			type
continue	for				import		return			var

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
append		bool			byte		cap			close		complex		complex64		complex128	uint16
copy		false			float32		float64		imag		int			int8			int16		uint32
int32		int64			iota		len			make		new			nil				panic		uint64
print		println			real		recover		string		true		uint			uint8		uintptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。
程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
程序中可能会使用到这些标点符号：.、,、;、: 和 …。


## Go 语言的空格
Go 语言中变量的声明必须使用空格隔开，如：
var age int
在变量与运算符间加入空格，程序看起来更加美观，如：
fruit = apples + oranges

*/

package main

import "fmt"

func caculate() {

	const a = 10
	const b = 20

	var c int
	c = a + b

	fmt.Println(a + b + c)

}

/* 占位符

普通占位符
占位符     说明                           举例                   输出
%v      相应值的默认格式。            Printf("%v", people)   {zhangsan}，
%+v     打印结构体时，会添加字段名     Printf("%+v", people)  {Name:zhangsan}
%#v     相应值的Go语法表示            Printf("#v", people)   main.Human{Name:"zhangsan"}
%T      相应值的类型的Go语法表示       Printf("%T", people)   main.Human
%%      字面上的百分号，并非值的占位符  Printf("%%")            %


----------------------------------------------------
布尔占位符
占位符       说明                举例                     输出
%t          true 或 false。     Printf("%t", true)       true


----------------------------------------------------
整数占位符
占位符     说明                                  举例                       输出
%b      二进制表示                             Printf("%b", 5)             101
%c      相应Unicode码点所表示的字符              Printf("%c", 0x4E2D)        中
%d      十进制表示                             Printf("%d", 0x12)          18
%o      八进制表示                             Printf("%d", 10)            12
%q      单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D)        '中'
%x      十六进制表示，字母形式为小写 a-f         Printf("%x", 13)             d
%X      十六进制表示，字母形式为大写 A-F         Printf("%x", 13)             D
%U      Unicode格式：U+1234，等同于 "U+%04X"   Printf("%U", 0x4E2D)         U+4E2D


----------------------------------------------------
浮点数和复数的组成部分（实部和虚部）
占位符     说明                              举例            输出
%b      无小数部分的，指数为二的幂的科学计数法，
        与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
%e      科学计数法，例如 -1234.456e+78        Printf("%e", 10.2)     1.020000e+01
%E      科学计数法，例如 -1234.456E+78        Printf("%e", 10.2)     1.020000E+01
%f      有小数点而无指数，例如 123.456        Printf("%f", 10.2)     10.200000
%g      根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出 Printf("%g", 10.20)   10.2
%G      根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出 Printf("%G", 10.20+2i) (10.2+2i)


----------------------------------------------------
字符串与字节切片
占位符     说明                              举例                           输出
%s      输出字符串表示（string类型或[]byte)   Printf("%s", []byte("Go语言"))  Go语言
%q      双引号围绕的字符串，由Go语法安全地转义  Printf("%q", "Go语言")         "Go语言"
%x      十六进制，小写字母，每字节两个字符      Printf("%x", "golang")         676f6c616e67
%X      十六进制，大写字母，每字节两个字符      Printf("%X", "golang")         676F6C616E67


----------------------------------------------------
指针
占位符         说明                      举例                             输出
%p      十六进制表示，前缀 0x          Printf("%p", &people)             0x4f57f0


----------------------------------------------------
其它标记
占位符      说明                             举例          输出
+      总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。
                                           Printf("%+q", "中文")  "\u4e2d\u6587"
-      在右侧而非左侧填充空格（左对齐该区域）
#      备用格式：为八进制添加前导 0（%#o）      Printf("%#U", '中')      U+4E2D
       为十六进制添加前导 0x（%#x）或 0X（%#X），为 %p（%#p）去掉前导 0x；
       如果可能的话，%q（%#q）会打印原始 （即反引号围绕的）字符串；
       如果是可打印字符，%U（%#U）会写出该字符的
       Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
' '    (空格)为数值中省略的正负号留出空白（% d）；
       以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
0      填充前导的0而非空格；对于数字，这会将填充移到正负号之后

*/
