// 1.Go 语言结构
/*
Go 语言的基础组成有以下几个部分：
	包声明
	引入包
	函数
	变量
	语句 & 表达式
	注释
*/

package main

// package main 定义了包名。
// 你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import "fmt"

// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），
// fmt 包实现了格式化 IO（输入/输出）的函数。
// 导入包后,可以使用格式 <包名>.<FuncName>调用函数
// 如果未调用包中的任何东西, 那么编译会出错,这样包中了尽可能少的导入无用的包, 因为导入的包会在程序运行时被链接,这样消耗性能

func main() {
	// 下一行 func main() 是程序开始执行的函数。只有package名为main的包可以包含main函数!
	// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
	// 一个可执行程序 有且仅有 一个main包

	/* main函数中打印字符串*/
	fmt.Println("Hello,世界!")
	caculate()

	// fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	// 使用 fmt.Print("hello, world\n") 可以得到相同的结果。
	// Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

	importFunc()
	testDef()

	getNumber()

	testForLoop()
	testLABEL()
	testArray()
	testSiceUsage()

	testResliceAppendCopy()

	testMap()

	testForRange()

	testSortMap()

	testInvoke()

	testStruct()

	testMethod()

	testValuePoint()

	testInterface()

	testSuperCollectionInterface()
}

// 标识符大小写区别:
// 大写:

// 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
// 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
// 小写:
// 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
