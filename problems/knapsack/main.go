package main

import (
	"fmt"
)

func knapsack(items [][]int, capacity int) int {
	n, w := len(items)+1, capacity+1
	table := make([][]int, n)
	for j := 0; j < n; j++ {
		table[j] = make([]int, w)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < w; j++ {
			if items[i-1][1] > j {
				table[i][j] = table[i-1][j]
			} else {
				k := table[i-1][j-(items[i-1][1])] + items[i-1][0]
				m := table[i-1][j]

				table[i][j] = max(k, m)
			}
		}
	}

	return table[n-1][w-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	items := [][]int{
		{5, 4},
		{10, 8},
		{3, 3},
		{2, 5},
		{3, 2},
	}
	capacity := 10

	fmt.Println(knapsack(items, capacity))
}
