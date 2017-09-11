package main

import "fmt"

func testVerifyKV() {

	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1] // OK pattern
	if !ok {
		m[2] = make(map[int]string)
	}

	m[2][1] = "GOOD"
	a, ok = m[2][1]
	fmt.Println(a, ok)
}
