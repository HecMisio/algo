package main

import (
	"algo/pkg/calculate"
	"fmt"
)

func main() {
	var exp = "(1 + 2*3 + (3-4) * 5) / 2" // expire 1 2 3 * + 3 4 - 5 * + 2 /
	post, err := calculate.Infix2Postfix(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("infix: %s\npostfix: %s\n", exp, post)

	ret, err := calculate.Calculate(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("calculate: %s = %v", exp, ret)
}
