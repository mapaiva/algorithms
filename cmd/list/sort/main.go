package main

import (
	"algorithms/list"
	"fmt"
)

func main() {
	s := &list.SortList{}

	items := []int{7, 3, 5, 2}
	for _, i := range items {
		s.Add(i)
		fmt.Println(s)
	}

	fmt.Println(s.Median())
}
