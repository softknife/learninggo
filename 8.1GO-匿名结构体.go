package main

import "fmt"

func testAnonamousStruct() {

	// 内部定义匿名结构, 直接使用指针接收
	a := &struct {
		Name string
		Age  int
	}{
		Name: "Eric",
		Age:  20,
	}

	fmt.Println(a)

}

type person1 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func testAnonamousStructEmbedded() {

	a := person1{Name: "eric-in", Age: 12}
	a.Contact.City = "xi'an"
	// a.Phone = "15721988888" // go中匿名嵌套时,不能这么写!
	a.Contact.Phone = "15721988888"

	fmt.Println(a)

}
