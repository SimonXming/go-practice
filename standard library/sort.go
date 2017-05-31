package main

import (
	"fmt"
	"sort"
)

type IntSlice2 []int

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	// s 类型转化 []int => IntSlice2
	var c IntSlice2 = IntSlice2(s)
	fmt.Println(c)

	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
	// [1 2 3 4 5 6]
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// [6 5 4 3 2 1]
}
