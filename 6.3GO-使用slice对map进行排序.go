/**

 */

package main

import (
	"fmt"
	"sort"
)

func testSortMap() {

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}

	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}

	sort.Ints(s)
	fmt.Println(s)

	mapV := make([]string, len(s))
	for i, v := range s {
		mapV[i] = m[v]
	}

	fmt.Println("sort result:", mapV)

	/*
		[1 2 3 4 5]
		sort result: [a b c d e]
	*/
}
