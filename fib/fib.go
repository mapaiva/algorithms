package fib

import "math/big"

// FibonacciN returns the n fibonacci number using recursion only solution: Complexity is O(2ˆ2).
func FibonacciN(n int) int {
	var result int

	if n == 1 || n == 2 {
		result = 1
	} else {
		result = FibonacciN(n-1) + FibonacciN(n-2)
	}

	return result
}

// FibonacciMemo returns the n fibonacci number using memoized solution: Complexity is O(n).
func FibonacciMemo(n int, memo []int) int {
	var result int

	if memo[n] != 0 {
		return memo[n]
	} else if n == 1 || n == 2 {
		result = 1
	} else {
		result = FibonacciMemo(n-1, memo) + FibonacciMemo(n-2, memo)
	}

	memo[n] = result

	return result
}

// FibonacciBottomUp returns the n fibonacci number using bottom up solution: Complexity is O(n)
func FibonacciBottomUp(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	bottomUp := make([]int, n+1)
	bottomUp[0] = 0
	bottomUp[1] = 1
	bottomUp[2] = 1

	for i := 3; i <= n; i++ {
		bottomUp[i] = bottomUp[i-1] + bottomUp[i-2]
	}

	return bottomUp[n]
}

// FibonacciBottomUpBig returns the n fibonacci number using bottom up solution and big Ints:
// Complexity is O(n).
func FibonacciBottomUpBig(n int) *big.Int {
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := 1; i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	return n1
}
