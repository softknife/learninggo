package main

// 这种导入
import "fmt"

// import "os"

// // 这样导入
// import (
// 	"strconv"
// 	"sort"
// 	"strings"
// )

// 注意,当前如果导入的包没有使用,vscode会帮我们直接删除无用的包

// 导入包别名
// import std "sort"

func importFunc() {

	var array = [...]int{1, 2, 3, 4}

	s1 := array[1:4]
	// s1 = std.Sort(s1)

	fmt.Println(array)
	fmt.Println(s1)

	fmt.Println("打印")
}
