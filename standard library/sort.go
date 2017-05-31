package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
	// [1 2 3 4 5 6]
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// [6 5 4 3 2 1]
}
