package main

import (
	"fmt"
)

func knapsack(items [][]int, capacity int) int {
	n, w := len(items)+1, capacity+1
	table := make([][]int, n)
	for j := 0; j < n; j++ { // O(n) onde n = len(items)+1 => O(n)
		table[j] = make([]int, w)
	}

	for i := 1; i < n; i++ { // O(n)
		for j := 0; j < w; j++ { // O(w) onde w = len(capacity)+1 => O(w)
			if weight(items, i-1) > j {
				table[i][j] = table[i-1][j]
			} else {
				k := table[i-1][j-(weight(items, i-1))] + value(items, i-1)
				m := table[i-1][j]

				table[i][j] = max(k, m)
			}
		}
	}

	// O(n) + [ O(n) * O(w) ] = O(n) + O(n*w) = O(n+n*w) = O(n*w)

	return table[n-1][w-1]
}

func weight(items [][]int, i int) int {
	return items[i][1]
}

func value(items [][]int, i int) int {
	return items[i][0]
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
