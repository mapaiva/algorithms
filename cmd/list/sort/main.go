package main

import (
	"algorithms/list"
	"fmt"
)

func main() {
	s := &list.SortList{}

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range items {
		s.Add(i)
		fmt.Println(s.Median())
	}
}
