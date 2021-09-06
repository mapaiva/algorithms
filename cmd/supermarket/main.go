package main

import (
	"fmt"
	"sort"
)

// Dado um conjunto de números, encontre os tres maiores valores dentre eles e se eles forem maior que 50% da soma total de todos os números printe EITA, senaão printe UFA.
//
// Complexity: O(n+n*log(n)).
func supermarketMood(supermarketItems []float64) {
	sum := 0.0
	for _, i := range supermarketItems { // O(n) being n = size(supermarketItems)
		sum += i
	}

	sort.Float64s(supermarketItems) // O(n*log(n)) being n = size(supermarketItems) performing a quicksort

	l := len(supermarketItems)
	item1 := supermarketItems[l-1]
	item2 := supermarketItems[l-2]
	item3 := supermarketItems[l-3]

	sumLast3 := item1 + item2 + item3

	if sumLast3/sum > 0.5 {
		fmt.Println("EITA")
	} else {
		fmt.Println("UFA")
	}
}

func main() {
	supermarketMood([]float64{
		15.23,
		21.3,
		12.45,
		4.5,
		2.99,
		10.0,
	})
}
