package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	num, err := strconv.ParseInt(arg1, 0, 64)
	if err != nil {
		panic(err)
	}

	r := big.NewInt(num)

	fmt.Println(factorial(r))
}

func factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		fmt.Println("n = ", n)
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}
