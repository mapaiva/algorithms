package main

import (
	"algorithms/fib"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Println(fib.FibonacciBottomUpBig(n))
}
